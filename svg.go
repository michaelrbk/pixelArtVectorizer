package main

import (
	"fmt"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/yourbasic/graph"
)

func genSVG(pixels [][]Pixel, graph graph.Mutable) {

	height := len(pixels) - 1
	width := len(pixels[0]) - 1
	scale := 50
	f, _ := os.Create("Vector.svg")
	canvas := svg.New(f)
	p := pixels[0][0]
	canvas.Start(width, height)

	//Print all pixel squares
	for y := 0; y <= width; y++ {
		for x := 0; x <= height; x++ {
			p = pixels[x][y]
			canvas.Rect(x, y, scale, scale, "fill:"+hexColor(p))
		}
	}

	canvas.End()

}

// hexColor returns an HTML hex-representation of c. The alpha channel is dropped
// and precision is truncated to 8 bits per channel
func hexColor(p Pixel) string {
	return fmt.Sprintf("#%.2x%.2x%.2x", p.R, p.G, p.B)
}
