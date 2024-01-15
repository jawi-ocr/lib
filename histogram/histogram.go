package histogram

import (
	"image"
)

func Get(img *image.Gray) []uint16 {
	bounds := img.Bounds()
	minRow := bounds.Min.Y
	maxRow := bounds.Max.Y
	minColumn := bounds.Min.X
	maxColumn := bounds.Max.X

	histogram := make([]uint16, maxRow)

	for row := minRow; row < maxRow; row++ {
		for column := minColumn; column < maxColumn; column++ {
			pixel := img.GrayAt(column, row).Y
			// 8-bit grayscale color
			// 0 = black
			// 255 = white
			if(pixel == 0) {
				histogram[row]++
			}
		}
	}

	return histogram
}

func GetWord(img *image.Gray) []uint16 {
	bounds := img.Bounds()
	minRow := bounds.Min.X
	maxRow := bounds.Max.X
	minColumn := bounds.Min.Y
	maxColumn := bounds.Max.Y

	histogram := make([]uint16, maxRow)

	for row := minRow; row < maxRow; row++ {
		for column := minColumn; column < maxColumn; column++ {
			pixel := img.GrayAt(row, column).Y
			// 8-bit grayscale color
			// 0 = black
			// 255 = white
			if(pixel == 0) {
				histogram[row]++
			}
		}
	}

	return histogram
}