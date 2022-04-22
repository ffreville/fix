package complete

import (
	"github.com/spf13/cobra"

	"sylr.dev/fix/pkg/dict"
	"sylr.dev/fix/pkg/utils"
)

func SecurityListRequestType(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return utils.PrettyOptionValues(dict.SecurityListRequestTypes), cobra.ShellCompDirectiveNoFileComp
}
