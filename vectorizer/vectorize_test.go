package vectorizer

import (
	"log"
	"os"
	"testing"
)

func TestVectorizeSuperMarioKart(t *testing.T) {

	Vectorize("./../testData/mario_kart.png")
	_, err := os.Open("./../results/3.pixelReshape.svg")
	if err != nil {
		log.Fatal(err)
	}
}
func TestVectorizeOverlapping(t *testing.T) {

	Vectorize("./../testData/background.png")
	_, err := os.Open("./../results/3.pixelReshape.svg")
	if err != nil {
		log.Fatal(err)
	}
}
func TestVectorizeMario(t *testing.T) {

	Vectorize("./../testData/mario_8bit.png")
	_, err := os.Open("./../results/3.pixelReshape.svg")
	if err != nil {
		log.Fatal(err)
	}
}
func TestVectorizeDolphin(t *testing.T) {

	Vectorize("./../testData/dolphin.png")
	_, err := os.Open("./../results/3.pixelReshape.svg")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCurveSize(t *testing.T) {
	g := NewGraph()
	p := (Pixel{})
	g.AddVertex(p)
	g.AddVertex(p)
	g.AddVertex(p)
	g.AddVertex(p)
	g.AddBoth(0, 1, Color{})
	g.AddBoth(1, 2, Color{})
	g.AddBoth(2, 3, Color{})
	n := curveSize(*g, 1, 2)
	if n != 5 {
		log.Fatal("Wrong Curve Size ", n)
	}
}
func TestCurveSizeLoop(t *testing.T) {
	g := NewGraph()
	p := (Pixel{})
	g.AddVertex(p)
	g.AddVertex(p)
	g.AddVertex(p)
	g.AddVertex(p)
	g.AddBoth(0, 1, Color{})
	g.AddBoth(1, 2, Color{})
	g.AddBoth(2, 3, Color{})
	g.AddBoth(3, 0, Color{})
	n := curveSize(*g, 2, 3)
	if n != 5 {
		log.Fatal("Wrong Curve Size ", n)
	}
}
