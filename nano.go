package blinkstick

import (
	"image/color"
	"strings"
	"time"

	"github.com/yesnault/hid"
)

// Nano represents a BlinkStick Nano https://www.blinkstick.com/products/blinkstick-nano
type Nano struct {
	usbDevice *usbDevice
}

func (nano Nano) getUSBDevice() *usbDevice {
	return nano.usbDevice
}

// ListFilter used for filter List Device
func (nano Nano) ListFilter(hid *hid.DeviceInfo) (bool, Blinkstick) {
	contains := strings.HasPrefix(hid.Product, "BlinkStick Nano")
	return contains, Nano{usbDevice: &usbDevice{DeviceInfo: hid}}
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

// Blink blink color for all led on current Blinkstick nano
func (nano Nano) Blink(color color.Color, duration, times int) error {
	if err := SetBlinkOnLed(nano, color, 0, duration, times); err != nil {
		return err
	}
	return SetBlinkOnLed(nano, color, 0, duration, times)
}

// SetColorTop set color for led on top on current Blinkstick nano
func (nano Nano) SetColorTop(color color.Color) error {
	if err := SetColorOnLed(nano, color, 0); err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)
	return nil
}

// SetColorBottom set color for bottom on top on current Blinkstick nano
func (nano Nano) SetColorBottom(color color.Color) error {
	if err := SetColorOnLed(nano, color, 1); err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)
	return nil
}
