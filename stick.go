package blinkstick

import (
	"image/color"
	"time"

	"github.com/yesnault/hid"
)

// Strip represents a BlinkStrip Strip https://www.blinkstick.com/products/blinkstick-strip
type Strip struct {
	usbDevice usbDevice
}

// ListFilter used for filter List Device
func (strip Strip) ListFilter(hid *hid.DeviceInfo) (bool, Blinkstick) {
	return hid.Product == "BlinkStick", Strip{usbDevice: usbDevice{DeviceInfo: hid}}
}

// GetDeviceInfo returns device info
func (strip Strip) GetDeviceInfo() *hid.DeviceInfo {
	return strip.usbDevice.DeviceInfo
}

// List returns blinkstick strip
func (strip Strip) List() []Blinkstick {
	return List(strip.ListFilter)
}

// SetColor set color for all led on current Blinkstick strip
func (strip Strip) SetColor(color color.Color) error {
	for index := 0; index < 8; index++ {
		if err := strip.usbDevice.setColor(byte(index), color); err != nil {
			return err
		}
		time.Sleep(1 * time.Millisecond)
	}
	return nil
}
