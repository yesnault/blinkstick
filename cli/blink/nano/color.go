package nano

import (
	"image/color"

	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
	"github.com/yesnault/blinkstick/cli/blink/internal"
)

var gcolor string
var top string
var bottom string

func init() {
	cmdNanoColor.PersistentFlags().StringVarP(&gcolor, "color", "", "", "Color for top and bottom led")
	cmdNanoColor.PersistentFlags().StringVarP(&top, "top", "", "", "Color for top led")
	cmdNanoColor.PersistentFlags().StringVarP(&bottom, "bottom", "", "", "Color for botton led")
}

var cmdNanoColor = &cobra.Command{
	Use:   "color",
	Short: "Color a blinkstick nano",
	Run: func(cmd *cobra.Command, args []string) {

		var colorColor, colorTop, colorBottom color.Color
		var err error

		if gcolor != "" {
			colorColor, err = blinkstick.GetColor(gcolor)
			internal.Check(err)
		}
		if top != "" {
			colorTop, err = blinkstick.GetColor(top)
			internal.Check(err)
		}
		if bottom != "" {
			colorBottom, err = blinkstick.GetColor(bottom)
			internal.Check(err)
		}

		b := blinkstick.Nano{}
		devices := b.List()
		for _, d := range devices {
			if gcolor != "" {
				internal.Check(d.SetColor(colorColor))
			}
			if top != "" {
				internal.Check(d.(blinkstick.Nano).SetColorTop(colorTop))
			}
			if bottom != "" {
				internal.Check(d.(blinkstick.Nano).SetColorBottom(colorBottom))
			}
		}

	},
}
