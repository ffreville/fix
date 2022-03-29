package complete

import (
	"github.com/spf13/cobra"

	"sylr.dev/fix/pkg/dict"
	"sylr.dev/fix/pkg/utils"
)

func OrderSide(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return utils.PrettyOptionValues(dict.OrderSidesReversed), cobra.ShellCompDirectiveNoFileComp
}

func OrderType(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return utils.PrettyOptionValues(dict.OrderTypesReversed), cobra.ShellCompDirectiveNoFileComp
}

func OrderTimeInForce(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return utils.PrettyOptionValues(dict.OrderTimeInForcesReversed), cobra.ShellCompDirectiveNoFileComp
}

func OrderPartyIDSource(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return utils.PrettyOptionValues(dict.PartyIDSourcesReversed), cobra.ShellCompDirectiveNoFileComp
}

func OrderPartyIDRole(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return utils.PrettyOptionValues(dict.PartyRolesReversed), cobra.ShellCompDirectiveNoFileComp
}

func OrderOriginationRole(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return utils.PrettyOptionValues(dict.OrderOriginationsReversed), cobra.ShellCompDirectiveNoFileComp
}
