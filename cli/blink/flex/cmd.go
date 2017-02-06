package flex

import (
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdFlexColor)
	Cmd.AddCommand(cmdFlexList)
}

// Cmd color
var Cmd = &cobra.Command{
	Use:   "flex",
	Short: "flex <command>",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
