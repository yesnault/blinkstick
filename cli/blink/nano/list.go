package nano

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
)

var cmdNanoList = &cobra.Command{
	Use:   "list",
	Short: "List all blinkstick nano",
	Run: func(cmd *cobra.Command, args []string) {

		b := blinkstick.Nano{}
		devices := b.List()
		for _, d := range devices {
			fmt.Printf("%+v \n", d.GetDeviceInfo())
		}

	},
}
