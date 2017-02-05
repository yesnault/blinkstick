package strip

import (
	"image/color"

	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
	"github.com/yesnault/blinkstick/cli/blink/internal"
)

var (
	serial     string
	gcolor     string
	leds       []int
	brightness int
	blink      int
	duration   int
)

func init() {
	cmdStripColor.PersistentFlags().StringVarP(&gcolor, "color", "", "", "Color for top and bottom led")
	cmdStripColor.PersistentFlags().IntVarP(&brightness, "brightness", "", 10, "Limit the brightness of the color 0..100")
	cmdStripColor.PersistentFlags().IntSliceVarP(&leds, "led", "", []int{}, "Led to manipulate: 0..7. If unspecified, action will be performed on all leds")
	cmdStripColor.PersistentFlags().StringVarP(&gcolor, "serial", "", "", "Select device by serial number. If unspecified, action will be performed on all BlinkSticks Nano")
	cmdStripColor.PersistentFlags().IntVarP(&blink, "blink", "", 0, "Blink LED --blink <ntimes> : --blink 1 for one time blinking")
	cmdStripColor.PersistentFlags().IntVarP(&duration, "duration", "", 100, "Set duration of transition in milliseconds (use with --blink)")
}

var cmdStripColor = &cobra.Command{
	Use:   "color",
	Short: "Color a blinkstick strip",
	Long: `Color a blinkstick strip:

Set the same color for all leds with 50% brightness :
  blink strip color --color orange --brightness 50

Color led 0 and 7
 blink strip color red --led 0 --led 7

Turn off light:
  blink strip color --color black

	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 || args[0] == "" {
			cmd.Help()
			return
		}
		gcolor := args[0]

		var colorColor color.Color
		var err error

		if gcolor != "" {
			colorColor, err = blinkstick.GetColor(gcolor, brightness)
			internal.Check(err)
		}

		b := blinkstick.Strip{}
		devices := b.List()
		for _, d := range devices {
			if serial == "" || d.GetDeviceInfo().SerialNumber == serial {
				if gcolor != "" {
					if len(leds) == 0 {
						if blink > 0 {
							internal.Check(d.Blink(colorColor, duration, blink))
						} else {
							internal.Check(d.SetColor(colorColor))
						}
					}
					for _, index := range leds {
						if index < 0 || index > 7 {
							continue
						}
						if blink > 0 {
							internal.Check(blinkstick.SetBlinkOnLed(d, colorColor, index, duration, blink))
						} else {
							internal.Check(blinkstick.SetColorOnLed(d, colorColor, index))
						}

					}
				}
			}
		}

	},
}
