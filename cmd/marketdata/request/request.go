package marketdatarequest

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"

	"sylr.dev/fix/config"
	"sylr.dev/fix/pkg/cli/complete"
	"sylr.dev/fix/pkg/dict"
	"sylr.dev/fix/pkg/errors"
	"sylr.dev/fix/pkg/initiator"
	"sylr.dev/fix/pkg/initiator/application"
	"sylr.dev/fix/pkg/utils"
)

var (
	optionTypes    []string
	optionSymbols  []string
	optionFullSnap bool
	optionSubType  string
	optionMDReqID  string
)

var MarketDataRequestCmd = &cobra.Command{
	Use:               "request",
	Short:             "Send a MarketDataRequest FIX message",
	Long:              "Send a MarketDataRequest FIX Message after initiating a session with a FIX acceptor.",
	Args:              cobra.ExactArgs(0),
	ValidArgsFunction: cobra.NoFileCompletions,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := Validate(cmd, args)
		if err != nil {
			return err
		}

		if cmd.HasParent() {
			parent := cmd.Parent()
			if parent.PersistentPreRunE != nil {
				return parent.PersistentPreRunE(cmd, args)
			}
		}
		return nil
	},
	RunE: Execute,
}

func init() {
	MarketDataRequestCmd.Flags().StringArrayVar(&optionSymbols, "symbol", []string{}, "Symbols")
	MarketDataRequestCmd.Flags().StringArrayVar(&optionTypes, "type", []string{"bid", "offer"}, "Order type (offer, bid, trade)")
	MarketDataRequestCmd.Flags().StringVar(&optionSubType, "sub-typ", "snapshot", "Subscription type")
	MarketDataRequestCmd.Flags().StringVar(&optionMDReqID, "id", "", "MarketDataRequest id (uuid autogenerated if not given)")
	utils.AddBothBoolFlags(MarketDataRequestCmd.Flags(), &optionFullSnap, "full", "", true, "Ask full refresh update")

	MarketDataRequestCmd.RegisterFlagCompletionFunc("symbol", cobra.NoFileCompletions)
	MarketDataRequestCmd.RegisterFlagCompletionFunc("type", complete.MDEntryTypes)
	MarketDataRequestCmd.RegisterFlagCompletionFunc("sub-type", complete.SubscriptionRequestTypes)
}

func Validate(cmd *cobra.Command, args []string) error {
	err := utils.ReconcileBoolFlags(cmd.Flags())
	if err != nil {
		return err
	}

	if len(optionSymbols) == 0 {
		return errors.OptionsNoSymbolGiven
	}

	if len(optionTypes) == 0 {
		return errors.OptionsNoTypeGiven
	}

	for _, t := range optionTypes {
		if _, ok := dict.MDEntryTypes[strings.ToUpper(t)]; !ok {
			return fmt.Errorf("%w: unkonwn type `%s`", errors.Options, t)
		}
	}

	if _, ok := dict.SubscriptionRequestTypes[strings.ToUpper(optionSubType)]; !ok {
		return fmt.Errorf("%w: unkonwn subscription type `%s`", errors.Options, optionSubType)
	}

	if len(optionMDReqID) == 0 {
		uid := uuid.New()
		optionMDReqID = uid.String()
	}

	return nil
}

func Execute(cmd *cobra.Command, args []string) error {
	options := config.GetOptions()
	logger := config.GetLogger()

	context, err := config.GetCurrentContext()
	if err != nil {
		return err
	}

	sessions, err := context.GetSessions()
	if err != nil {
		return err
	}

	ctxInitiator, err := context.GetInitiator()
	if err != nil {
		return err
	}

	session := sessions[0]
	transportDict, appDict, err := session.GetFIXDictionaries()
	if err != nil {
		return err
	}

	settings, err := context.ToQuickFixInitiatorSettings()
	if err != nil {
		return err
	}

	app := application.NewMarketDataRequest()
	app.Logger = logger
	app.Settings = settings
	app.TransportDataDictionary = transportDict
	app.AppDataDictionary = appDict

	var quickfixLogger *zerolog.Logger
	if options.QuickFixLogging {
		quickfixLogger = logger
	}

	init, err := initiator.Initiate(app, settings, quickfixLogger)
	if err != nil {
		return err
	}

	// Start session
	err = init.Start()
	if err != nil {
		return err
	}

	// Choose right timeout cli option > config > default value (5s)
	var timeout time.Duration
	if options.Timeout != time.Duration(0) {
		timeout = options.Timeout
	} else if ctxInitiator.SocketTimeout != time.Duration(0) {
		timeout = ctxInitiator.SocketTimeout
	} else {
		timeout = 5 * time.Second
	}

	// Wait for session connection
	select {
	case <-time.After(timeout):
		return errors.ConnectionTimeout
	case _, ok := <-app.Connected:
		if !ok {
			return errors.FixLogout
		}
	}

	// Prepare securitylist
	securitylist, err := buildMessage(*session)
	if err != nil {
		return err
	}

	// Send the order
	err = quickfix.Send(securitylist)
	if err != nil {
		return err
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

LOOP:
	for {
		select {
		case signal := <-interrupt:
			logger.Debug().Msgf("Received signal: %s", signal)

			app.Stop()
			init.Stop()

			break LOOP
		case _, ok := <-app.FromAppChan:
			if !ok {
				break LOOP
			}
		}
	}

	return nil
}

func buildMessage(session config.Session) (quickfix.Messagable, error) {
	mdReqID := field.NewMDReqID(optionMDReqID)
	subReqType := field.NewSubscriptionRequestType(dict.SubscriptionRequestTypes[strings.ToUpper(optionSubType)])
	marketDepth := field.NewMarketDepth(0)

	var updateType enum.MDUpdateType
	if optionFullSnap {
		updateType = enum.MDUpdateType_FULL_REFRESH
	} else {
		updateType = enum.MDUpdateType_INCREMENTAL_REFRESH
	}

	// Message
	message := quickfix.NewMessage()
	header := fixt11.NewHeader(&message.Header)

	switch session.BeginString {
	case quickfix.BeginStringFIXT11:
		switch session.DefaultApplVerID {
		case "FIX.5.0SP2":
			header.Set(field.NewMsgType(enum.MsgType_MARKET_DATA_REQUEST))
			message.Body.Set(mdReqID)
			message.Body.Set(subReqType)
			message.Body.Set(marketDepth)
			message.Body.Set(field.NewMDUpdateType(updateType))

			entryTypes := quickfix.NewRepeatingGroup(
				tag.NoMDEntryTypes,
				quickfix.GroupTemplate{
					quickfix.GroupElement(tag.MDEntryType),
				},
			)
			for _, t := range optionTypes {
				entryTypes.Add().Set(field.NewMDEntryType(dict.MDEntryTypes[strings.ToUpper(t)]))
			}
			message.Body.SetGroup(entryTypes)

			relatedSym := quickfix.NewRepeatingGroup(
				tag.NoRelatedSym,
				quickfix.GroupTemplate{
					quickfix.GroupElement(tag.Symbol),
				},
			)
			for _, sym := range optionSymbols {
				relatedSym.Add().Set(field.NewSymbol(sym))
			}
			message.Body.SetGroup(relatedSym)
		default:
			return nil, errors.FixVersionNotImplemented
		}
	default:
		return nil, errors.FixVersionNotImplemented
	}

	utils.QuickFixMessagePartSetString(&message.Header, session.TargetCompID, field.NewTargetCompID)
	utils.QuickFixMessagePartSetString(&message.Header, session.TargetSubID, field.NewTargetSubID)
	utils.QuickFixMessagePartSetString(&message.Header, session.SenderCompID, field.NewSenderCompID)
	utils.QuickFixMessagePartSetString(&message.Header, session.SenderSubID, field.NewSenderSubID)

	return message, nil
}
