package color

import (
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdColorList)
}

// Cmd color
var Cmd = &cobra.Command{
	Use:   "color",
	Short: "Color list",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
