package blinkstick

import (
	"fmt"
	"image/color"

	"golang.org/x/image/colornames"
)

// GetColor returns a color from a name
func GetColor(s string, brightness int) (color.Color, error) {
	if color, ok := colornames.Map[s]; ok {
		return applyBrightness(color, brightness), nil
	}
	return nil, fmt.Errorf("Invalid color %s", s)
}

// ColorList returns all available colors
func ColorList() []string {
	return colornames.Names
}

// brightness : 0-100
func applyBrightness(c color.Color, brightness int) color.Color {
	r, g, b, _ := c.RGBA()
	if brightness > 100 {
		brightness = 100
	}
	if brightness < 0 {
		brightness = 0
	}

	// int(float(options.limit) / 100.0 * 255)
	rn := remapColor(uint8(r), brightness)
	gn := remapColor(uint8(g), brightness)
	bn := remapColor(uint8(b), brightness)

	return color.RGBA{R: rn, G: gn, B: bn}
}

func remap(value float64, leftMin, leftMax, rightMin, rightMax float64) uint8 {
	// Figure out how 'wide' each range is
	leftSpan := leftMax - leftMin
	rightSpan := rightMax - rightMin

	// Convert the left range into a 0-1 range (float)
	valueScaled := float64(value-leftMin) / float64(leftSpan)

	// Convert the 0-1 range into a value in the right range.
	return uint8(rightMin + (valueScaled * rightSpan))
}

func remapColor(value uint8, max int) uint8 {
	limit := float64(max) / 100 * 255
	r := remap(float64(value), 0, 255, 0, limit)
	return r
}
