package vectorizer

// solveAmbiguities
// analyse block of 2x2 vertex to delete the max amount of edges and also remover cross connections between these vertexes/pixels
// 00 10
// 01 11
func solveAmbiguities(pixels [][]Pixel, g *PixelArtGraph, genSVG bool) {

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

			//Crossing edges that we need to eliminate
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
						if curveSize(*g, p01.V, p10.V) <= curveSize(*g, p11.V, p00.V) {
							g.DeleteBoth(p01.V, p10.V)
						} else {
							g.DeleteBoth(p11.V, p00.V)
						}
					} else {

						//heuristic of overlapping pixels
						sumC1 := 0
						sumC2 := 0
						//start x and y with -4 positions to check 3 pixels in both ways
						xs := x - 4
						ys := y - 4
						c1 := p00.Color
						c2 := p01.Color

						for xs <= x+3 {
							for ys <= y+3 {
								if p00.Color == c1 {
									sumC1++
								} else if p00.Color == c2 {
									sumC2++
								}
								ys++
							}
							xs++
						}
						//the color in the largest amount represents the background and should be kept connected
						if sumC1 > sumC2 {
							g.DeleteBoth(p11.V, p00.V)
						} else {
							g.DeleteBoth(p01.V, p10.V)
						}
					}

				}
			}
		}
	}
	if genSVG {
		generateSVG(pixels, g, SvgConfig{"./results/2.solveAmbiguities", 50, true, true, true, false})
	}

}

//curveSize return the size of the 1 pixel line
func curveSize(g PixelArtGraph, vertexA int, vertexB int) int {
	size := 0
	hasEdge := true
	va := -1
	vb := -1
	if g.Degree(vertexA) == 2 || g.Degree(vertexB) == 2 {
		size++
		for hasEdge {

			hasEdge = false
			if g.Degree(vertexA) == 2 {
				size++

				g.Visit(vertexA, func(w int) (skip bool) {
					if w != va && w != vb && w != vertexA && w != vertexB {
						va = w
						hasEdge = true
					}

					return
				})

			}

			if g.Degree(vertexB) == 2 {
				size++

				g.Visit(vertexB, func(w int) (skip bool) {
					if w != vb && w != va && w != vertexA && w != vertexB {
						vb = w
						hasEdge = true
					}
					return
				})
			}

		}
	}
	return size
}
