package vectorizer

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

// SvgConfig configurations to SVG file
type SvgConfig struct {
	fileName     string
	scale        int
	withPixel    bool
	withVertices bool
	withEdges    bool
	withPoints   bool
}

func generateSVG(pixels [][]Pixel, g *Graph, config SvgConfig) {
	scale := config.scale
	height := len(pixels)
	width := len(pixels[0])

	f, _ := os.Create(config.fileName + ".svg")
	canvas := svg.New(f)
	p := pixels[0][0]
	canvas.Start(height*scale, width*scale)

	//Print all pixels squares
	if config.withPixel {
		for y := 0; y < width; y++ {
			for x := 0; x < height; x++ {
				p = pixels[x][y]
				canvas.Rect(x*scale, y*scale, scale, scale,
					"fill=\""+p.Color.hexColor()+"\" stroke=\"Black\" stroke-width=\"1\"")
			}
		}
	}

	//Print all Vertex center points
	if config.withVertices {
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
	}

	if config.withPoints {

		for y := 0; y < width; y++ {
			for x := 0; x < height; x++ {
				p = pixels[x][y]
				var sX []int
				var sY []int
				for z := 0; z < len(p.Points); z++ {
					sX = append(sX, p.Points[z].X)
					sY = append(sY, p.Points[z].Y)
				}
				if len(sX) > 0 {
					canvas.Polygon(sX, sY,
						"fill=\""+p.Color.hexColor()+"\" stroke=\"Black\" stroke-width=\"0.2\"")
				}
			}
		}
	}

	//Print all edges lines
	if config.withEdges {
		p1 := Pixel{}
		p2 := Pixel{}
		for v := 0; v < g.Order(); v++ {
			g.Visit(v, func(w int) (skip bool) {
				p1 = getPixelV(pixels, v)
				p2 = getPixelV(pixels, w)
				canvas.Line(p1.X*scale+scale/2, p1.Y*scale+scale/2,
					p2.X*scale+scale/2, p2.Y*scale+scale/2,
					"stroke=\"Blue\"", "stroke-width=\"0.8\"")
				return
			})

		}
	}
	canvas.End()
	f.Close()

}
