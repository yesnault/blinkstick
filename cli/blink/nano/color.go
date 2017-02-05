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
var brightness int

func init() {
	cmdNanoColor.PersistentFlags().StringVarP(&gcolor, "color", "", "", "Color for top and bottom led")
	cmdNanoColor.PersistentFlags().StringVarP(&top, "top", "", "", "Color for top led")
	cmdNanoColor.PersistentFlags().StringVarP(&bottom, "bottom", "", "", "Color for botton led")
	cmdNanoColor.PersistentFlags().IntVarP(&brightness, "brightness", "", 10, "Limit the brightness of the color 0..100")
}

var cmdNanoColor = &cobra.Command{
	Use:   "color",
	Short: "Color a blinkstick nano",
	Long: `Color a blinkstick nano:

Set the same color for both led:
  blink nano color --color orange

Set a color for bottom Led and another for top Led:
  blink nano color --bottom red --top green

Turn off light:
	blink nano color --color black

	`,
	Run: func(cmd *cobra.Command, args []string) {

		if gcolor == "" && top == "" && bottom == "" {
			cmd.Help()
			return
		}

		var colorColor, colorTop, colorBottom color.Color
		var err error

		if gcolor != "" {
			colorColor, err = blinkstick.GetColor(gcolor, brightness)
			internal.Check(err)
		}
		if top != "" {
			colorTop, err = blinkstick.GetColor(top, brightness)
			internal.Check(err)
		}
		if bottom != "" {
			colorBottom, err = blinkstick.GetColor(bottom, brightness)
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
