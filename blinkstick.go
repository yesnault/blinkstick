package blinkstick

import (
	"fmt"
	"image/color"

	"github.com/boombuler/hid"
)

// Version of Blinkstick
// One Line for this, used by release.sh script
// Keep "const Version on one line"
const Version = "0.0.1"

// VendorID blinkstick
const VendorID = 0x20a0

// ProductID blinkstick
const ProductID = 0x41e5

// USBDevice ...
type USBDevice struct {
	DeviceInfo *hid.DeviceInfo
	Device     *hid.Device
}

// Blinkstick represents a blinkstick device
type Blinkstick interface {
	List() []Blinkstick
	SetColor(color.Color) error
	GetDeviceInfo() *hid.DeviceInfo
	ListFilter(hid *hid.DeviceInfo) (bool, Blinkstick)
}

// SetColor set color
func (usbDevice *USBDevice) setColor(index byte, c color.Color) error {
	if usbDevice.Device == nil {
		if err := usbDevice.Open(); err != nil {
			return err
		}
	}
	r, g, b, _ := c.RGBA()
	d := *usbDevice.Device
	return d.WriteFeature([]byte{0x05, 0x00, index, byte(r >> 8), byte(g >> 8), byte(b >> 8)})
}

// Open open a device
func (usbDevice *USBDevice) Open() error {
	device, err := usbDevice.DeviceInfo.Open()
	if err != nil {
		return fmt.Errorf("Error while opening device: %s", err)
	}
	usbDevice.Device = &device
	return nil
}

// ListFilter is used to filter device on List
type ListFilter func(*hid.DeviceInfo) (bool, Blinkstick)

// List gets all blinkstick device
func List(opts ...ListFilter) []Blinkstick {
	out := []Blinkstick{}

	if len(opts) == 0 {
		opts = append(opts, Nano{}.ListFilter)
	}

	for di := range hid.Devices() {
		if di.VendorId == VendorID && di.ProductId == ProductID {
			for _, o := range opts {
				if toKeep, blinkstick := o(di); toKeep {
					out = append(out, blinkstick)
				}
			}

		}
	}
	return out
}
