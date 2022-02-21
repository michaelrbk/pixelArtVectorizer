package vectorizer

import (
	"fmt"
	"image"
)

// Pixel struct
type Pixel struct {
	V      int // index in graph
	Color  Color
	X      int
	Y      int
	Points []Point
}

//Point struct
type Point struct {
	X int
	Y int
}

// Color of Pixel
type Color struct {
	R int
	G int
	B int
	A int
}

// hexColor returns an HTML hex-representation of c. The alpha channel is dropped
// and precision is truncated to 8 bits per channel
func (c Color) hexColor() string {
	return fmt.Sprintf("#%.2x%.2x%.2x", c.R, c.G, c.B)
}

// Get the bi-dimensional pixel array
func getPixels(img image.Image) ([][]Pixel, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	vCount := 0
	for x := 0; x < width; x++ {
		var col []Pixel
		for y := 0; y < height; y++ {
			col = append(col, Pixel{vCount, rgbaToColor(img.At(x, y).RGBA()), x, y, []Point{}})
			vCount++
		}
		pixels = append(pixels, col)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToColor(r uint32, g uint32, b uint32, a uint32) Color {
	return Color{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
