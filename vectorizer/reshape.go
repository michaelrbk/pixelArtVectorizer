package vectorizer

//reshape all the pixel in the original PixelArt
//here we analyse every pixel checking the surrounding connections in a 3x3 window around it
func reshape(pixels [][]Pixel, g *PixelArtGraph, genSVG bool) {
	scale := 7

	//Visit all Vertices
	for v := 0; v < g.Order(); v++ {
		//Top Left of new pixel
		/*
			if has edge to tl
			   -2 +2
			   +2 -2
			if the top has edge to bl
				+2 +2
			else
				0 0
		*/
		x := getPixelV(pixels, v).X
		y := getPixelV(pixels, v).Y
		p := getPixel(pixels, x, y)

		tl := getPixel(pixels, x-1, y-1)
		t := getPixel(pixels, x, y-1)
		tr := getPixel(pixels, x+1, y-1)

		l := getPixel(pixels, x-1, y)
		r := getPixel(pixels, x+1, y)

		br := getPixel(pixels, x+1, y+1)
		b := getPixel(pixels, x, y+1)
		bl := getPixel(pixels, x-1, y+1)

		//Edge between center pixel and top left
		if g.Edge(p.V, tl.V) {
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale - 2, y*scale + 2})
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 2, y*scale - 2})
		} else if g.Edge(t.V, l.V) { //Edge between the top pixel and the left
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 2, y*scale + 2})
		} else {
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 0, y*scale + 0})
		}

		//Top Right of new pixel
		/*
			if has edge to tr
				+5 -2
				+9 +2
			if the top has edge br
				+5  +2
			else
				+7  +0
		*/
		if g.Edge(p.V, tr.V) { //Edge between center pixel and top right
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 5, y*scale - 2})
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 9, y*scale + 2})
		} else if g.Edge(t.V, r.V) { //Edge between top pixel and  pixel in the right
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 5, y*scale + 2})
		} else {
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 7, y*scale + 0})
		}

		//Bottom Right of new pixel
		/*
			if has edge to br
				  +9 +5
				  +5 +9
			if the top has edge to tr
				  +5  +5
			else
				  +7  +7
		*/
		if g.Edge(p.V, br.V) { //Edge between center pixel and bottom right
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 9, y*scale + 5})
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 5, y*scale + 9})
		} else if g.Edge(b.V, r.V) { //Edge between bottom pixel and the pixel at the right
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 5, y*scale + 5})
		} else {
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 7, y*scale + 7})
		}

		//Bottom Left of new pixel
		/*
			if has edge to bl
				   +2 +9
				   -2 +5
			if the top has edge to tl
				   +2 +5
			else
				   +0 +7
		*/
		if g.Edge(p.V, bl.V) { //Edge between center pixel and bottom left
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 2, y*scale + 9})
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale - 2, y*scale + 5})
		} else if g.Edge(b.V, l.V) { //Edge between below pixel and bottom pixel
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 2, y*scale + 5})
		} else {
			pixels[x][y].NewShapePoints = append(pixels[x][y].NewShapePoints, Point{x*scale + 0, y*scale + 7})
		}

	}
	if genSVG {
		generateSVG(pixels, g, SvgConfig{"./results/3.pixelReshape", 7, false, false, false, true})
	}
}
