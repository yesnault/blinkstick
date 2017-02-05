package internal

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"

	"github.com/yesnault/blinkstick"
)

// Check checks error, if != nil, throw panic
func Check(e error) {
	if e != nil {
		log.Fatalf("%s", e)
	}
}

// Exit func display an error message on stderr and exit 1
func Exit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

// DisplayDevices display list of devices, formatted
func DisplayDevices(devices []blinkstick.Blinkstick) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Manufacturer", "Description", "Serial"})

	for _, d := range devices {
		v := d.GetDeviceInfo()
		table.Append([]string{
			v.Manufacturer,
			v.Product,
			v.SerialNumber,
		})
	}
	table.Render() // Send output

}
