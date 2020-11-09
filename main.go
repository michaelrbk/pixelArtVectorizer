package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	imagePath := os.Args[1:]
	if len(imagePath) == 0 {
		fmt.Println("No image parameter received")
		return
	}
	_, err := readImage(imagePath[0])
	if err != nil {
		fmt.Println("Error reading image", err)
		return
	}

	// newGraph(img)
	// solveAmbiguities()
	// reshapePixelCell()
	// drawNewGraphEdges()
	// createNewCurves()
	fmt.Println("imagePath", imagePath)
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