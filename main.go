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
	g := newGraph(pixels)
	genSVG(pixels, *g)
	// solveAmbiguities()
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

func newGraph(pixels [][]Pixel) *graph.Mutable {

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
