package strip

import (
	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
	"github.com/yesnault/blinkstick/cli/blink/internal"
)

var cmdStripList = &cobra.Command{
	Use:   "list",
	Short: "List all blinkstick strip",
	Run: func(cmd *cobra.Command, args []string) {
		b := blinkstick.Strip{}
		internal.DisplayDevices(b.List())
	},
}
