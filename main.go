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
	g := genGraph(pixels)
	genSVG(pixels, *g)
	// solveAmbiguities(pixels, *g)
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
func genGraph(pixels [][]Pixel) *graph.Mutable {

	width := len(pixels)
	height := len(pixels[0])
	g := graph.New((height + 1) * (width + 1))
	xc := 0
	yc := 0
	p := pixels[0][0]
	pCompare := p

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			p = pixels[x][y]

			xc = x - 1
			yc = y - 1
			if xc > 0 && yc > 0 && xc < height && yc < width {
				pCompare = pixels[xc][yc]
				if p == pCompare {
					//2d to 1d array
					g.AddBoth(a2dTo1d(x, y, height), a2dTo1d(xc, yc, height))
				}
			}

			xc = x
			yc = y - 1
			if xc > 0 && yc > 0 && xc < height && yc < width {
				pCompare = pixels[xc][yc]
				if p == pCompare {
					//2d to 1d array
					g.AddBoth(a2dTo1d(x, y, height), a2dTo1d(xc, yc, height))
				}
			}

			xc = x + 1
			yc = y - 1
			if xc > 0 && yc > 0 && xc < height && yc < width {
				pCompare = pixels[xc][yc]
				if p == pCompare {
					// if x <= xc || y <= yc {
					//2d to 1d array
					g.AddBoth(a2dTo1d(x, y, height), a2dTo1d(xc, yc, height))

				}
			}
			xc = x - 1
			yc = y
			p = pixels[x][y]
			if xc > 0 && yc > 0 && xc < height && yc < width {
				pCompare = pixels[xc][yc]
				if p == pCompare {
					//2d to 1d array
					g.AddBoth(a2dTo1d(x, y, height), a2dTo1d(xc, yc, height))
				}
			}

		}
	}
	return g
}

// solveAmbiguities
// analyse block of 2x2 vertex to delete the max amount of edges and alse remover cross conections between this vertexes/pixels
// 00 10
// 01 11
func solveAmbiguities(pixels [][]Pixel, g graph.Mutable) {

	width := len(pixels)
	height := len(pixels[0])

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p00 := a2dTo1d(x, y, width)
			c00 := pixels[x][y]

			p10 := a2dTo1d(x+1, y, width)
			//c10 := pixels[x+1][y]

			p01 := a2dTo1d(x, y+1, width)
			c01 := pixels[x][y+1]

			p11 := a2dTo1d(x+1, y+1, width)
			//c11 := pixels[x+1][y+1]

			//Crossing edges that we need to elimnate
			if g.Edge(p11, p00) && g.Edge(p01, p10) {
				// Ambiguity
				if c00 == c01 { //Same color
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

}
func curveSize(g *graph.Mutable, p00 int, p01 int, p10 int, p11 int) int {
	return 0
}
