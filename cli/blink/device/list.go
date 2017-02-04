package device

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
)

func init() {
}

var cmdDeviceList = &cobra.Command{
	Use:   "list",
	Short: "List all blinkstick device: blinkstick device list",
	Run: func(cmd *cobra.Command, args []string) {

		devices := blinkstick.List()

		// TODO format list
		for _, d := range devices {
			fmt.Printf("%+v \n", d.GetDeviceInfo())
		}

	},
}
