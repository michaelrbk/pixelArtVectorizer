package main

import (
	"fmt"
	"os"
)

func main() {
	imagePath := os.Args[1:]
	if len(imagePath) == 0 {
		fmt.Println("No image parameter received")
		return
	}
	img, err := readImage(imagePath[0])
	if err != nil {
		fmt.Println("Error reading image", err)
		return
	}

	pixels, err := getPixels(img)
	if err != nil {
		fmt.Println("Error converting to bi-dimensional array", err)
		return
	}
	//bi-dimensional pixel array
	g := genGraph(pixels, true)

	solveAmbiguities(pixels, &g, true)
	reshape(pixels, &g, true)
	// drawNewGraphEdges()
	// createNewCurves()

}

/*
newGraph generate a graph from pixel art with connections
mario_8bit.png example
	Dimensions 19 x 18
	width 19
	height 18
*/
func genGraph(pixels [][]Pixel, genSVG bool) Graph {

	width := len(pixels)
	height := len(pixels[0])

	g := NewGraph(width * height)
	xc := 0
	yc := 0
	p := Pixel{}
	pc := p //Comparison pixel

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			p = pixels[x][y]

			xc = x - 1

			yc = y - 1
			pc = getPixel(pixels, xc, yc)
			if p.Color == pc.Color {
				g.AddBoth(p.V, pc.V, p.Color)
			}

			xc = x
			yc = y - 1
			pc = getPixel(pixels, xc, yc)
			if p.Color == pc.Color {
				g.AddBoth(p.V, pc.V, p.Color)
			}

			xc = x + 1
			yc = y - 1
			pc = getPixel(pixels, xc, yc)
			if p.Color == pc.Color {
				g.AddBoth(p.V, pc.V, p.Color)

			}
			xc = x - 1
			yc = y

			pc = getPixel(pixels, xc, yc)
			if p.Color == pc.Color {
				g.AddBoth(p.V, pc.V, p.Color)
			}
		}
	}
	if genSVG {

		generateSVG(pixels, g, SvgConfig{"./results/0.source", 50, true, false, false, false})
		generateSVG(pixels, g, SvgConfig{"./results/1.genGraph", 50, true, true, true, false})
	}
	return *g

}

func getPixel(pixels [][]Pixel, x int, y int) Pixel {
	width := len(pixels)
	height := len(pixels[0])

	if x >= 0 && y >= 0 && x < width && y < height {
		return pixels[x][y]
	}
	return Pixel{V: -1}

}

//getpixelv from Vertice index
func getPixelV(pixels [][]Pixel, v int) Pixel {
	width := len(pixels)
	height := len(pixels[0])
	p := Pixel{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p = getPixel(pixels, x, y)
			if p.V == v {
				return p
			}
		}
	}
	return p
}
