package color

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
)

var cmdColorList = &cobra.Command{
	Use:   "list",
	Short: "List all colors",
	Run: func(cmd *cobra.Command, args []string) {

		colors := blinkstick.ColorList()

		for _, c := range colors {
			fmt.Printf("%+v \n", c)
		}

	},
}
