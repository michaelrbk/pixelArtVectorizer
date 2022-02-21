package main

import (
	"github.com/michaelrbk/pixelArtVectorizer/vectorizer"
	"os"
)

func main() {
	imagePath := os.Args[1]
	vectorizer.Vectorize(imagePath)
}
