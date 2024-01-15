package monochrome

import (
	"image"
	"image/color"
)

// Monochrome is an in-memory image whose At method returns Pixel values.
type Monochrome struct {
	// Pix holds the image's pixels, as gray values. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func NewMonochrome(r image.Rectangle) *Monochrome {
	w, h := r.Dx(), r.Dy()
	stride := (w + 7) / 8
	pix := make([]uint8, stride*h)
	return &Monochrome{pix, stride, r}
}

func (p *Monochrome) ColorModel() color.Model { 
	return MonochromeModel
}

func (p *Monochrome) Bounds() image.Rectangle { 
	return p.Rect 
}

func (p *Monochrome) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return Pixel(White)
	}
	i := p.PixOffset(x, y)
	return Pixel(p.Pix[i]&(1<<uint(x%8)) != 0)
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Monochrome) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*1
}

func (p *Monochrome) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	if MonochromeModel.Convert(c).(Pixel) == Black {
		p.Pix[i] |= 1 << uint(x%8)
	} else {
		p.Pix[i] &^= 1 << uint(x%8)
	}
}

// TODO: Coefficient can be selected
func ConvertImage(img image.Image, threshold uint8) *image.Gray {
	bounds := img.Bounds()
	monochromeImage := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			// Get the pixel color at (x,y)
			pixelColor := img.At(x,y)

			// Get the RGB value
			r, g, b, _ := pixelColor.RGBA()

			// Convert RGB value to uin8
			r8 := r>>8
			g8 := g>>8
			b8 := b>>8

			var binaryColor color.Gray
			// grayValue := uint8(0.2126*float32(r) + 0.7152*float32(g) + 0.0722*float32(b))
			grayValue := 0.2126*float32(r8) + 0.7152*float32(g8) + 0.0722*float32(b8)
			if grayValue > float32(threshold) {
				// Set White
				binaryColor.Y = 255
			} else {
				// Set Black
				binaryColor.Y = 0
			}

			monochromeImage.SetGray(x, y, binaryColor)
		}
	}


	return monochromeImage
}