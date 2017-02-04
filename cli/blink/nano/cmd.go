package nano

import (
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdNanoColor)
	Cmd.AddCommand(cmdNanoList)
}

// Cmd color
var Cmd = &cobra.Command{
	Use:   "nano",
	Short: "nano <command>",
}
