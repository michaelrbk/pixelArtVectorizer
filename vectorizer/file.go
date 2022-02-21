package vectorizer

import (
	"image"
	"image/png"
	"os"
)

func readImage(file string) (image.Image, error) {
	// Read image from file that already exists
	existingImageFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		return nil, err
	}

	return loadedImage, nil
}
