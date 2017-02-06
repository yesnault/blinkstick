package flex

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
	blink      bool
	duration   int
	repeats    int
)

func init() {
	cmdFlexColor.PersistentFlags().StringVarP(&gcolor, "color", "", "", "Color for top and bottom led: blink flex color red")
	cmdFlexColor.PersistentFlags().IntVarP(&brightness, "brightness", "", 10, "Limit the brightness of the color 0..100")
	cmdFlexColor.PersistentFlags().IntSliceVarP(&leds, "led", "", []int{}, "Led to manipulate: 0..7. If unspecified, action will be performed on all leds")
	cmdFlexColor.PersistentFlags().StringVarP(&gcolor, "serial", "", "", "Select device by serial number. If unspecified, action will be performed on all BlinkSticks Nano")
	cmdFlexColor.PersistentFlags().BoolVarP(&blink, "blink", "", false, "Blink LED")
	cmdFlexColor.PersistentFlags().IntVarP(&duration, "duration", "", 100, "Set duration of transition in milliseconds (use with --blink)")
	cmdFlexColor.PersistentFlags().IntVarP(&repeats, "repeats", "", 10, "Number of repetitions (use with --blink)")
}

var cmdFlexColor = &cobra.Command{
	Use:   "color",
	Short: "Color a blinkstick flex",
	Long: `Color a blinkstick flex:

Set the same color for all leds with 50% brightness :
  blink flex color orange --brightness 50

Color led 0 and 7
 blink flex color red --led 0 --led 7

Turn off light:
  blink flex color black

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

		b := blinkstick.Flex{}
		devices := b.List()
		for _, d := range devices {
			if serial == "" || d.GetDeviceInfo().SerialNumber == serial {
				if gcolor != "" {
					if len(leds) == 0 {
						if blink {
							internal.Check(d.Blink(colorColor, duration, repeats))
						} else {
							internal.Check(d.SetColor(colorColor))
						}
					}
					for _, index := range leds {
						if index < 0 || index > 32 {
							continue
						}
						if blink {
							internal.Check(blinkstick.SetBlinkOnLed(d, colorColor, index, duration, repeats))
						} else {
							internal.Check(blinkstick.SetColorOnLed(d, colorColor, index))
						}
					}
				}
			}
		}

	},
}
