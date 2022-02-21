package vectorizer

import "fmt"

func Vectorize(imagePath string) {
	if len(imagePath) == 0 {
		fmt.Println("No image parameter received")
		return
	}
	img, err := readImage(imagePath)
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
	// newGraphEdges(pixels, &g, true)
	// createNewCurves()
}
