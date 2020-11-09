package main

import (
	"image"
)

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

// Get the bi-dimensional pixel array
func getPixels(img image.Image) ([][]Pixel, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for x := 0; x < width; x++ {
		var col []Pixel
		for y := 0; y < height; y++ {
			col = append(col, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, col)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// a2dTo1d return the position in a 1d Array
func a2dTo1d(x int, y int, width int) int {
	return x + (y * width)
}

// a1dTo2d return the position in a 2d Array
func a1dTo2d(index int, width int) (x int, y int) {
	return index % width, index / width
}
