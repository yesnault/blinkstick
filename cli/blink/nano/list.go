package nano

import (
	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
	"github.com/yesnault/blinkstick/cli/blink/internal"
)

var cmdNanoList = &cobra.Command{
	Use:   "list",
	Short: "List all blinkstick nano",
	Run: func(cmd *cobra.Command, args []string) {
		b := blinkstick.Nano{}
		internal.DisplayDevices(b.List())
	},
}
