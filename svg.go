package main

import (
	"fmt"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/yourbasic/graph"
)

func genSVG(pixels [][]Pixel, graph graph.Mutable) {
	scale := 50
	height := len(pixels) - 1
	width := len(pixels[0]) - 1

	f, _ := os.Create("Vector.svg")
	canvas := svg.New(f)
	p := pixels[0][0]
	canvas.Start(width*scale+1, height*scale+1)
	//Print all vertices/pixels squares
	for y := 0; y <= width; y++ {
		for x := 0; x <= height; x++ {
			p = pixels[x][y]
			canvas.Rect(x*scale, y*scale, scale, scale, "fill:"+hexColor(p))
		}
	}
	//Print all edges lines
	x1, y1, x2, y2 := 0, 0, 0, 0
	for v := 0; v < graph.Order(); v++ {
		graph.Visit(v, func(w int, c int64) (skip bool) {
			x1, y1 = a1dTo2d(v, width)
			x2, y2 = a1dTo2d(w, width)
			return
		})
		//stroke="Blue" stroke-width="0.8"
		canvas.Line(x1*scale+scale/2, y1*scale+scale/2, x2*scale+scale/2, y2*scale+scale/2, "stroke=\"Blue\"", "stroke-width=\"0.8\"")
	}
	canvas.End()

}

// hexColor returns an HTML hex-representation of c. The alpha channel is dropped
// and precision is truncated to 8 bits per channel
func hexColor(p Pixel) string {
	return fmt.Sprintf("#%.2x%.2x%.2x", p.R, p.G, p.B)
}
