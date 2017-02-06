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
	blink      bool
	duration   int
	repeats    int
)

func init() {
	cmdNanoColor.PersistentFlags().StringVarP(&top, "top", "", "", "Color for top led: blink nano color red")
	cmdNanoColor.PersistentFlags().StringVarP(&bottom, "bottom", "", "", "Color for botton led")
	cmdNanoColor.PersistentFlags().IntVarP(&brightness, "brightness", "", 1, "Limit the brightness of the color 0..100")
	cmdNanoColor.PersistentFlags().StringVarP(&serial, "serial", "", "", "Select device by serial number. If unspecified, action will be performed on all BlinkSticks Strip")
	cmdNanoColor.PersistentFlags().BoolVarP(&blink, "blink", "", false, "Blink LED")
	cmdNanoColor.PersistentFlags().IntVarP(&duration, "duration", "", 100, "Set duration of transition in milliseconds (use with --blink)")
	cmdNanoColor.PersistentFlags().IntVarP(&repeats, "repeats", "", 10, "Number of repetitions (use with --blink)")
}

var cmdNanoColor = &cobra.Command{
	Use:   "color",
	Short: "Color a blinkstick nano: blink nano color [<color>] [--brightness=n] [--top=<color>] [--bottom=<color>] [--serial=s] [--duration=n] [--repeats=n] [--blink]",
	Long: `Color a blinkstick nano:

Set the same color for both led with 50% brightness :
  blink nano color orange --brightness 50

Set a color for bottom Led and another for top Led:
  blink nano color --bottom red --top green

Examples:
 blink nano color --top purple --brightness 1
 blink nano color --bottom red --brightness 100
 blink nano color green --brightness 12
 blink nano color --serial BS008173-3.0 --duration=500 --repeats=10 --brightness 30 --blink --bottom red

Turn off light:
  blink nano color black

	`,
	Run: func(cmd *cobra.Command, args []string) {

		if (len(args) == 0 || args[0] == "") && top == "" && bottom == "" {
			cmd.Help()
			return
		}
		gcolor := ""
		if len(args) > 0 {
			gcolor = args[0]
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
			if serial == "" || d.GetDeviceInfo().SerialNumber == serial {
				if gcolor != "" {
					if blink {
						internal.Check(d.Blink(colorColor, duration, repeats))
					} else {
						internal.Check(d.SetColor(colorColor))
					}
				}
				if blink {
					if top != "" {
						internal.Check(blinkstick.SetBlinkOnLed(d, colorTop, 0, duration, repeats))
					}
					if bottom != "" {
						internal.Check(blinkstick.SetBlinkOnLed(d, colorBottom, 1, duration, repeats))
					}
				} else {
					if top != "" {
						internal.Check(d.(blinkstick.Nano).SetColorTop(colorTop))
					}
					if bottom != "" {
						internal.Check(d.(blinkstick.Nano).SetColorBottom(colorBottom))
					}
				}

			}
		}

	},
}
