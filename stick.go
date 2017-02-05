package blinkstick

import (
	"image/color"

	"github.com/yesnault/hid"
)

// Strip represents a BlinkStrip Strip https://www.blinkstick.com/products/blinkstick-strip
type Strip struct {
	usbDevice *usbDevice
}

func (strip Strip) getUSBDevice() *usbDevice {
	return strip.usbDevice
}

// ListFilter used for filter List Device
func (strip Strip) ListFilter(hid *hid.DeviceInfo) (bool, Blinkstick) {
	return hid.Product == "BlinkStick", Strip{usbDevice: &usbDevice{DeviceInfo: hid}}
}

// GetDeviceInfo returns device info
func (strip Strip) GetDeviceInfo() *hid.DeviceInfo {
	return strip.usbDevice.DeviceInfo
}

// List returns blinkstick strip
func (strip Strip) List() []Blinkstick {
	return List(strip.ListFilter)
}

// Blink blink color for all led on current Blinkstick strip
func (strip Strip) Blink(color color.Color, duration, times int) error {
	for index := 0; index < 8; index++ {
		if err := SetBlinkOnLed(strip, color, index, duration, times); err != nil {
			return err
		}
	}
	return nil
}

// SetColor set color for all led on current Blinkstick strip
func (strip Strip) SetColor(color color.Color) error {
	for index := 0; index < 8; index++ {
		SetColorOnLed(strip, color, index)
	}
	return nil
}
