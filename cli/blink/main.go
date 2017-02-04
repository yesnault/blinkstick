package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/yesnault/blinkstick/cli/blink/color"
	"github.com/yesnault/blinkstick/cli/blink/device"
	"github.com/yesnault/blinkstick/cli/blink/nano"
	"github.com/yesnault/blinkstick/cli/blink/update"
	"github.com/yesnault/blinkstick/cli/blink/version"
)

var rootCmd = &cobra.Command{
	Use:   "blink",
	Short: "Blink - Command Line for Blinkstick",
}

func init() {
	flags := rootCmd.Flags()

	flags.String("log-level", "", "Log Level : debug, info or warn")
	viper.BindPFlag("log_level", flags.Lookup("log-level"))
}

func main() {
	addCommands()

	switch viper.GetString("log_level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Err:%s", err)
	}
}

//AddCommands adds child commands to the root command rootCmd.
func addCommands() {
	rootCmd.AddCommand(color.Cmd)
	rootCmd.AddCommand(device.Cmd)
	rootCmd.AddCommand(nano.Cmd)
	rootCmd.AddCommand(update.Cmd)
	rootCmd.AddCommand(version.Cmd)

}
