package strip

import (
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdStripColor)
	Cmd.AddCommand(cmdStripList)
}

// Cmd color
var Cmd = &cobra.Command{
	Use:   "strip",
	Short: "strip <command>",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
