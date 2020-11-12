package main

func reshape(pixels [][]Pixel, g *Graph, genSVG bool) {
	// width := len(pixels)
	// height := len(pixels[0])
	scale := 7

	//Visit all Vertices
	for v := 0; v < g.Order(); v++ {
		//TOP LEFT
		/*
			if has tl
			   -2 +2
			   +2 -2
			if the top has bl
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
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale - 2, y*scale + 2})
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 2, y*scale - 2})
		} else if g.Edge(t.V, l.V) { //Edge between the top pixel and the left
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 2, y*scale + 2})
		} else {
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 0, y*scale + 0})
		}

		//TOP RIGHT
		/*
			if has tr
				+5 -2
				+9 +2
			if the top has br
				+5  +2
			else
				+7  +0
		*/

		if g.Edge(p.V, tr.V) { //Edge between center pixel and top right
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 5, y*scale - 2})
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 9, y*scale + 2})
		} else if g.Edge(t.V, r.V) { //Edge between top pixel and  pixel in the right
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 5, y*scale + 2})
		} else {
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 7, y*scale + 0})
		}

		//botton righ
		/*
			if has br
				  +9 +5
				  +5 +9
			if the top has tr
				  +5  +5
			else
				  +7  +7
		*/

		if g.Edge(p.V, br.V) { //Edge between center pixel and below right
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 9, y*scale + 5})
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 5, y*scale + 9})
		} else if g.Edge(b.V, r.V) { //Edge between below pixel and the pixel to the righ
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 5, y*scale + 5})
		} else {
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 7, y*scale + 7})
		}

		//BOTTON LEFT
		/*
			if has bl
				   +2 +9
				   -2 +5
			if the top has tl
				   +2 +5
			else
				   +0 +7
		*/
		if g.Edge(p.V, bl.V) { //Edge between center pixel and below left
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 2, y*scale + 9})
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale - 2, y*scale + 5})
		} else if g.Edge(b.V, l.V) { //Edge between below pixel and left pixel
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 2, y*scale + 5})
		} else {
			pixels[x][y].Points = append(pixels[x][y].Points, Point{x*scale + 0, y*scale + 7})
		}

	}
	if genSVG {
		generateSVG(pixels, g, SvgConfig{"./results/3.pixelReshape", 7, false, false, false, true})
	}
}
