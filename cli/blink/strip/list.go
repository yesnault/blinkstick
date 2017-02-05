package strip

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
)

var cmdStripList = &cobra.Command{
	Use:   "list",
	Short: "List all blinkstick strip",
	Run: func(cmd *cobra.Command, args []string) {

		b := blinkstick.Strip{}
		devices := b.List()
		for _, d := range devices {
			fmt.Printf("%+v \n", d.GetDeviceInfo())
		}

	},
}
