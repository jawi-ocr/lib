package monochrome

import (
	"image/color"
)

// Pixel represents a 1-bit color.
type Pixel bool

// Monochrome color is either a black or white.
// 0 = White
// 1 = Black
const (
	Black Pixel = true
	White Pixel = false
)

// MonochromeModel for the color of black and white type.
var MonochromeModel color.Model = color.ModelFunc(monochromeModel)

// MonochromeColor represents a 1-bit color.
type MonochromeColor struct {
	Pixel
}

func (c MonochromeColor) RGBA() (r, g, b, a uint32) {
	return 0xffff, 0xffff, 0xffff, 0xffff
}

func (c Pixel) RGBA() (r, g, b, a uint32) {
	if c == Black {
		return 0, 0, 0, 0xffff
	}
	return 0xffff, 0xffff, 0xffff, 0xffff
}


// https://github.com/ev3go/ev3dev/blob/a5fda5c6a492269e01b184046ed42dc4a1dfe8c9/fb/mono.go

func monochromeModel(c color.Color) color.Color {
	if _, ok := c.(Pixel); ok {
		return c
	}
	r, g, b, _ := c.RGBA()

	// Reference: 
	// https://cs.opensource.google/go/go/+/refs/tags/go1.21.0:src/image/image.go
	// https://cs.opensource.google/go/go/+/refs/tags/go1.21.0:src/image/color/color.go

	// These coefficients (the fractions 0.299, 0.587 and 0.114) are the same
	// as those given by the JFIF specification and used by func RGBToYCbCr in
	// ycbcr.go.
	//
	// Note that 19595 + 38470 + 7471 equals 65536.
	//
	// The 24 is 16 + 8. The 16 is the same as used in RGBToYCbCr. The 8 is
	// because the return value is 8 bit color, not 16 bit color.
	// y := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
	// uint8(y)

	// Reference: https://en.wikipedia.org/wiki/Grayscale
	// Coefficient(s) option:
	// #1: y := 0.299*r + 0.587*g + 0.114*b
	// #2: y := 0.2126*r + 0.7152*g + 0.0722*b
	// #3: y : = uint8((r + g + b) / 3 >> 8)

	y := uint8(0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b))

	return Pixel(y < 0x8)
}