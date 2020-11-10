package main

import (
	"fmt"
	"os"

	"github.com/yourbasic/graph"

	svg "github.com/ajstarks/svgo"
)

func generateSVG(pixels [][]Pixel, g graph.Mutable, fileName string) {
	scale := 50
	height := len(pixels)
	width := len(pixels[0])
	a2dSize := len(pixels[0])

	f, _ := os.Create(fileName + ".svg")
	canvas := svg.New(f)
	p := pixels[0][0]
	canvas.Start(width*scale, height*scale)
	//Print all pixels squares
	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			p = pixels[x][y]
			canvas.Rect(x*scale, y*scale, scale, scale, "fill=\""+hexColor(p)+"\" stroke=\"Black\" stroke-width=\"1\"")
		}
	}
	//Print all Vertex points

	r := scale / 10 //radius
	if scale < 10 {
		r = 1
	}

	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			p = pixels[x][y]

			canvas.Circle(x*scale+scale/2, y*scale+scale/2, r, "fill=\"blue\" stroke=\"Black\" stroke-width=\"1\"")
		}
	}
	//Print all edges lines
	x1, y1, x2, y2 := 0, 0, 0, 0
	for v := 0; v < g.Order(); v++ {
		g.Visit(v, func(w int, c int64) (skip bool) {
			x1, y1 = a1dTo2d(v, a2dSize)
			x2, y2 = a1dTo2d(w, a2dSize)
			return
		})
		//stroke="Blue" stroke-width="0.8"
		canvas.Line(x1*scale+scale/2, y1*scale+scale/2, x2*scale+scale/2, y2*scale+scale/2, "stroke=\"Blue\"", "stroke-width=\"0.8\"")
	}
	canvas.End()
	f.Close()

}

// hexColor returns an HTML hex-representation of c. The alpha channel is dropped
// and precision is truncated to 8 bits per channel
func hexColor(p Pixel) string {
	return fmt.Sprintf("#%.2x%.2x%.2x", p.R, p.G, p.B)
}
