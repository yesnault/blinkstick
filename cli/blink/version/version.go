package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yesnault/blinkstick"
)

// Cmd version
var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Display Version of blink: blink version",
	Long:    `blink version`,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version blink: %s\n", blinkstick.Version)
	},
}
