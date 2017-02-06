package flex

import (
	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
	"github.com/yesnault/blinkstick/cli/blink/internal"
)

var cmdFlexList = &cobra.Command{
	Use:   "list",
	Short: "List all blinkstick flex",
	Run: func(cmd *cobra.Command, args []string) {
		b := blinkstick.Flex{}
		internal.DisplayDevices(b.List())
	},
}
