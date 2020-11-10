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
	a2dSize := len(pixels[0]) - 1
	g := graph.New(16)
	xc := 0
	yc := 0
	p := Pixel{}
	pc := p //Comparison pixel

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			p = pixels[x][y]

			xc = x - 1

			yc = y - 1
			pc = getPixel(pixels, x-1, y-1)
			if p == pc {
				g.AddBoth(a2dTo1d(x, y, a2dSize), a2dTo1d(xc, yc, a2dSize))
			}

			xc = x
			yc = y - 1
			pc = getPixel(pixels, x, y-1)
			if p == pc {
				g.AddBoth(a2dTo1d(x, y, a2dSize), a2dTo1d(xc, yc, a2dSize))
			}

			xc = x + 1
			yc = y - 1
			pc = getPixel(pixels, x+1, y-1)
			if p == pc {
				g.AddBoth(a2dTo1d(x, y, a2dSize), a2dTo1d(xc, yc, a2dSize))

			}
			xc = x - 1
			yc = y

			pc = getPixel(pixels, x-1, y)
			if p == pc {
				g.AddBoth(a2dTo1d(x, y, a2dSize), a2dTo1d(xc, yc, a2dSize))
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

	//Vertex being analysed
	p00, p10 := 0, 0
	p01, p11 := 0, 0

	//Pixel Color
	c00, c01 := Pixel{}, Pixel{}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p00 = a2dTo1d(x, y, height)
			c00 = getPixel(pixels, x, y)

			p10 = a2dTo1d(x+1, y, height)

			p01 = a2dTo1d(x, y+1, height)
			c01 = getPixel(pixels, x, y+1)

			p11 = a2dTo1d(x+1, y+1, height)

			//Crossing edges that we need to elimnate
			if g.Edge(p11, p00) && g.Edge(p01, p10) {
				// Ambiguity
				if c00 == c01 && c00 != (Pixel{}) { //Same color and not empty pixel
					if g.Edge(p11, p00) {
						g.Delete(p11, p00)
					}
					if g.Edge(p01, p10) {
						g.Delete(p01, p10)
					}

					// island heuristic
				} else if g.Degree(p11) == 1 || g.Degree(p00) == 1 { //if is alone need to keep connected
					if g.Edge(p01, p10) {
						g.Delete(p01, p10)
					}
				} else if g.Degree(p01) == 1 || g.Degree(p10) == 1 {
					if g.Edge(p11, p00) {
						g.Delete(p11, p00)
					}
				} else {
					// curve heuristic
					if g.Degree(p01) == 2 || g.Degree(p10) == 2 || g.Degree(p11) == 2 || g.Degree(p00) == 2 { // is part of a curve the bigger curve is keep connected
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
	return 0
}

func getPixel(pixels [][]Pixel, x int, y int) Pixel {
	width := len(pixels)
	height := len(pixels[0])

	if x >= 0 && y >= 0 && x < width && y < height {
		println(x, y)
		return pixels[x][y]
	}
	return Pixel{}

}
