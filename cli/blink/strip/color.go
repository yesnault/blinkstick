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
	brightness int
)

func init() {
	cmdStripColor.PersistentFlags().StringVarP(&gcolor, "color", "", "", "Color for top and bottom led")
	cmdStripColor.PersistentFlags().IntVarP(&brightness, "brightness", "", 10, "Limit the brightness of the color 0..100")
	cmdStripColor.PersistentFlags().StringVarP(&gcolor, "serial", "", "", "Select device by serial number. If unspecified, action will be performed on all BlinkSticks Nano")
}

var cmdStripColor = &cobra.Command{
	Use:   "color",
	Short: "Color a blinkstick strip",
	Long: `Color a blinkstick strip:

Set the same color for all leds with 50% brightness :
  blink strip color --color orange --brightness 50

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
					internal.Check(d.SetColor(colorColor))
				}
			}
		}

	},
}
