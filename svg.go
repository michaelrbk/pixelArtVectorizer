package main

import (
	"os"

	"github.com/yourbasic/graph"

	svg "github.com/ajstarks/svgo"
)

func generateSVG(pixels [][]Pixel, g graph.Mutable, fileName string) {
	scale := 50
	height := len(pixels)
	width := len(pixels[0])

	f, _ := os.Create(fileName + ".svg")
	canvas := svg.New(f)
	p := pixels[0][0]
	canvas.Start(width*scale, height*scale)

	//Print all pixels squares
	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			p = pixels[x][y]
			canvas.Rect(x*scale, y*scale, scale, scale,
				"fill=\""+p.Color.hexColor()+"\" stroke=\"Black\" stroke-width=\"1\"")
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
			canvas.Circle(x*scale+scale/2, y*scale+scale/2, r,
				"fill=\"blue\" stroke=\"Black\" stroke-width=\"1\"")
		}
	}

	//Print all edges lines
	p1 := Pixel{}
	p2 := Pixel{}
	for v := 0; v < g.Order(); v++ {
		g.Visit(v, func(w int, c int64) (skip bool) {
			p1 = getPixelV(pixels, v)
			p2 = getPixelV(pixels, w)
			canvas.Line(p1.X*scale+scale/2, p1.Y*scale+scale/2,
				p2.X*scale+scale/2, p2.Y*scale+scale/2,
				"stroke=\"Blue\"", "stroke-width=\"0.8\"")
			return
		})

	}
	canvas.End()
	f.Close()

}
