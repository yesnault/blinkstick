package blinkstick

import (
	"fmt"
	"image/color"

	"golang.org/x/image/colornames"
)

// GetColor returns a color from a name
func GetColor(s string) (color.Color, error) {
	if color, ok := colornames.Map[s]; ok {
		return color, nil
	}
	return nil, fmt.Errorf("Invalid color %s", s)
}

// ColorList returns all available colors
func ColorList() []string {
	return colornames.Names
}
