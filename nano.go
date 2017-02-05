package blinkstick

import (
	"image/color"
	"strings"
	"time"

	"github.com/boombuler/hid"
)

// Nano represents a BlinkStick Nano https://www.blinkstick.com/products/blinkstick-nano
type Nano struct {
	usbDevice usbDevice
}

// ListFilter used for filter List Device
func (nano Nano) ListFilter(hid *hid.DeviceInfo) (bool, Blinkstick) {
	contains := strings.HasPrefix(hid.Product, "BlinkStick Nano")
	return contains, Nano{usbDevice: usbDevice{DeviceInfo: hid}}
}

// GetDeviceInfo returns device info
func (nano Nano) GetDeviceInfo() *hid.DeviceInfo {
	return nano.usbDevice.DeviceInfo
}

// List returns blinkstick nano
func (nano Nano) List() []Blinkstick {
	return List(nano.ListFilter)
}

// SetColor set color for all led on current Blinkstick nano
func (nano Nano) SetColor(color color.Color) error {
	if err := nano.SetColorTop(color); err != nil {
		return err
	}
	return nano.SetColorBottom(color)
}

// SetColorTop set color for led on top on current Blinkstick nano
func (nano Nano) SetColorTop(color color.Color) error {
	if err := nano.usbDevice.setColor(0, color); err != nil {
		return err
	}
	time.Sleep(1 * time.Millisecond)
	return nil
}

// SetColorBottom set color for bottom on top on current Blinkstick nano
func (nano Nano) SetColorBottom(color color.Color) error {
	if err := nano.usbDevice.setColor(1, color); err != nil {
		return err
	}
	time.Sleep(1 * time.Millisecond)
	return nil
}
