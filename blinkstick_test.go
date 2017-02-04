package blinkstick

import "testing"

func TestList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: ("test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			List()
		})
	}
}
