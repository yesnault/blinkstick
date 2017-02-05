package nano

import (
	"image/color"

	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
	"github.com/yesnault/blinkstick/cli/blink/internal"
)

var (
	top        string
	bottom     string
	brightness int
	serial     string
)

func init() {
	cmdNanoColor.PersistentFlags().StringVarP(&top, "top", "", "", "Color for top led")
	cmdNanoColor.PersistentFlags().StringVarP(&bottom, "bottom", "", "", "Color for botton led")
	cmdNanoColor.PersistentFlags().IntVarP(&brightness, "brightness", "", 1, "Limit the brightness of the color 0..100")
	cmdNanoColor.PersistentFlags().StringVarP(&serial, "serial", "", "", "Select device by serial number. If unspecified, action will be performed on all BlinkSticks Strip")
}

var cmdNanoColor = &cobra.Command{
	Use:   "color",
	Short: "Color a blinkstick nano: blink nano color <color> [--brightness 1]",
	Long: `Color a blinkstick nano:

Set the same color for both led with 50% brightness :
  blink nano color --color orange --brightness 50

Set a color for bottom Led and another for top Led:
  blink nano color --bottom red --top green

Turn off light:
  blink nano color --color black

	`,
	Run: func(cmd *cobra.Command, args []string) {

		if (len(args) == 0 || args[0] == "") && top == "" && bottom == "" {
			cmd.Help()
			return
		}
		gcolor := args[0]

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
			if serial == "" || d.GetDeviceInfo().SerialNumber == serial {
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
		}

	},
}
