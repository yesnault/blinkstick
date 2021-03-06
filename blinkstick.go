package blinkstick

import (
	"fmt"
	"image/color"

	"github.com/yesnault/hid"
)

// Version of Blinkstick
// One Line for this, used by release.sh script
// Keep "const Version on one line"
const Version = "0.1.7"

// vendorID blinkstick
const vendorID = 0x20a0

// productID blinkstick
const productID = 0x41e5

// usbDevice ...
type usbDevice struct {
	DeviceInfo *hid.DeviceInfo
	Device     *hid.Device
}

// Blinkstick represents a blinkstick device
type Blinkstick interface {
	List() []Blinkstick
	SetColor(color.Color) error
	Blink(color.Color, int, int) error
	GetDeviceInfo() *hid.DeviceInfo
	ListFilter(hid *hid.DeviceInfo) (bool, Blinkstick)
	getUSBDevice() *usbDevice
}

// SetColor set color
func (usbDevice *usbDevice) setColor(c color.Color, index byte) error {
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
func (usbDevice *usbDevice) Open() error {
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
		opts = append(opts, Flex{}.ListFilter)
		opts = append(opts, Strip{}.ListFilter)
	}

	for di := range hid.Devices() {
		if di.VendorId == vendorID && di.ProductId == productID {
			//			fmt.Printf("di: %+v", di)
			for _, o := range opts {
				if toKeep, blinkstick := o(di); toKeep {
					out = append(out, blinkstick)
				}
			}

		}
	}
	return out
}
