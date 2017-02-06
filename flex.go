package blinkstick

import (
	"image/color"
	"strings"

	"github.com/yesnault/hid"
)

// Flex represents a BlinkFlex Flex https://www.blinkstick.com/products/blinkstick-flex
type Flex struct {
	usbDevice *usbDevice
}

func (flex Flex) getUSBDevice() *usbDevice {
	return flex.usbDevice
}

// ListFilter used for filter List Device
func (flex Flex) ListFilter(hid *hid.DeviceInfo) (bool, Blinkstick) {
	contains := strings.HasPrefix(hid.Product, "BlinkStick Flex")
	return contains, Flex{usbDevice: &usbDevice{DeviceInfo: hid}}
}

// GetDeviceInfo returns device info
func (flex Flex) GetDeviceInfo() *hid.DeviceInfo {
	return flex.usbDevice.DeviceInfo
}

// List returns blinkstick flex
func (flex Flex) List() []Blinkstick {
	return List(flex.ListFilter)
}

// Blink blink color for all led on current Blinkstick flex
func (flex Flex) Blink(color color.Color, duration, times int) error {
	for index := 0; index < 32; index++ {
		if err := SetBlinkOnLed(flex, color, index, duration, times); err != nil {
			return err
		}
	}
	return nil
}

// SetColor set color for all led on current Blinkstick flex
func (flex Flex) SetColor(color color.Color) error {
	for index := 0; index < 32; index++ {
		SetColorOnLed(flex, color, index)
	}
	return nil
}
