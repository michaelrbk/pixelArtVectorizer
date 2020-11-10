package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/yourbasic/graph"
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

	solveAmbiguities(pixels, *g, true)
	// reshapePixelCell()
	// drawNewGraphEdges()
	// createNewCurves()

}

func readImage(file string) (image.Image, error) {
	// Read image from file that already exists
	existingImageFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		return nil, err
	}

	fmt.Println(imageData)
	fmt.Println(imageType)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		return nil, err
	}
	fmt.Println(loadedImage)
	return loadedImage, nil
}

/*
newGraph generate a graph from pixel art with connections
mario_8bit.png example
	Dimensions 19 x 18
	width 19
	height 18
*/
func genGraph(pixels [][]Pixel, genSVG bool) *graph.Mutable {

	width := len(pixels)
	height := len(pixels[0])
	g := graph.New(width * height)
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
				g.AddBoth(p.V, pc.V)
			}

			xc = x
			yc = y - 1
			pc = getPixel(pixels, xc, yc)
			if p.Color == pc.Color {
				g.AddBoth(p.V, pc.V)
			}

			xc = x + 1
			yc = y - 1
			pc = getPixel(pixels, xc, yc)
			if p.Color == pc.Color {
				g.AddBoth(p.V, pc.V)

			}
			xc = x - 1
			yc = y

			pc = getPixel(pixels, xc, yc)
			if p.Color == pc.Color {
				g.AddBoth(p.V, pc.V)
			}
		}
	}
	if genSVG {
		generateSVG(pixels, *g, "./results/1.genGraph")
	}
	return g
}

// solveAmbiguities
// analyse block of 2x2 vertex to delete the max amount of edges and alse remover cross conections between this vertexes/pixels
// 00 10
// 01 11
func solveAmbiguities(pixels [][]Pixel, g graph.Mutable, genSVG bool) {

	width := len(pixels)
	height := len(pixels[0])

	//Pixel being analysed
	p00, p10 := Pixel{}, Pixel{}
	p01, p11 := Pixel{}, Pixel{}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p00 = getPixel(pixels, x, y)
			p10 = getPixel(pixels, x+1, y)
			p01 = getPixel(pixels, x, y+1)
			p11 = getPixel(pixels, x+1, y+1)

			//Crossing edges that we need to elimenate
			if g.Edge(p11.V, p00.V) && g.Edge(p01.V, p10.V) {
				// Ambiguity
				if p00.Color == p01.Color && p00.Color != (Color{}) { //Same color and not empty pixel
					if g.Edge(p11.V, p00.V) {
						g.DeleteBoth(p11.V, p00.V)
					}
					if g.Edge(p01.V, p10.V) {
						g.DeleteBoth(p01.V, p10.V)
					}

					// island heuristic
				} else if g.Degree(p11.V) == 1 || g.Degree(p00.V) == 1 { //if is alone need to keep connected
					if g.Edge(p01.V, p10.V) {
						g.DeleteBoth(p01.V, p10.V)
					}
				} else if g.Degree(p01.V) == 1 || g.Degree(p10.V) == 1 {
					if g.Edge(p11.V, p00.V) {
						g.DeleteBoth(p11.V, p00.V)
					}
				} else {
					// curve heuristic
					if g.Degree(p01.V) == 2 || g.Degree(p10.V) == 2 || g.Degree(p11.V) == 2 || g.Degree(p00.V) == 2 { // is part of a curve the bigger curve is keep connected
						// if curve size from edge1 <= edge2
						// 	remove edge1
						// else remove edge 2
					} else {
						//heuristic of overlapping pixels
					}
				}
			}
		}
	}
	if genSVG {
		generateSVG(pixels, g, "./results/2.solveAmbiguities")
	}

}
func curveSize(g graph.Mutable, p00 int, p01 int, p10 int, p11 int) int {
	// todo:
	return 0
}

func getPixel(pixels [][]Pixel, x int, y int) Pixel {
	width := len(pixels)
	height := len(pixels[0])

	if x >= 0 && y >= 0 && x < width && y < height {
		return pixels[x][y]
	}
	return Pixel{}

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
