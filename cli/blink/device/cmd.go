package device

import (
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdDeviceList)
}

// Cmd color
var Cmd = &cobra.Command{
	Use:   "device",
	Short: "device <command>",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
