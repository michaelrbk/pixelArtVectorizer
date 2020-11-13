package main

import (
	"log"
	"os"
	"testing"
)

func TestVectorizeMario(t *testing.T) {

	vectorize("./testData/mario_8bit.png")
	_, err := os.Open("./results/3.pixelReshape.svg")
	if err != nil {
		log.Fatal(err)
	}
}

func TestVectorizeSuperMarioKart(t *testing.T) {

	//vectorize("./testData/superMarioKart.png")
	vectorize("./testData/curves_loop.png")
	_, err := os.Open("./results/3.pixelReshape.svg")
	if err != nil {
		log.Fatal(err)
	}
}

func TestVectorizeDolphin(t *testing.T) {

	vectorize("./testData/dolphin.png")
	_, err := os.Open("./results/3.pixelReshape.svg")
	if err != nil {
		log.Fatal(err)
	}
}
