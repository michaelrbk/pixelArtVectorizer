package vectorizer

import (
	"strconv"
)

const initialMapSize = 4

// PixelArtGraph with pixel information
type PixelArtGraph struct {
	edges  map[int]map[int]Color
	vertex map[int]Pixel
}

// NewGraph constructs a new graph
func NewGraph() *PixelArtGraph {
	return &PixelArtGraph{edges: make(map[int]map[int]Color), vertex: make(map[int]Pixel)}
}

// Order returns the number of vertices in the graph.
func (g *PixelArtGraph) Order() int {
	return len(g.edges)
}

// Edge tells if there is an edge from v to w.
func (g *PixelArtGraph) Edge(v, w int) bool {
	if v < 0 || v >= g.Order() {
		return false
	}
	if w < 0 || w >= g.Order() {
		return false
	}
	_, ok := g.edges[v][w]

	return ok
}

// AddVertex and return its index
func (g *PixelArtGraph) AddVertex(p Pixel) int {
	v := len(g.vertex)
	g.vertex[v] = p

	g.edges[v] = make(map[int]Color)
	return v
}

// Add inserts a directed edge from v to w with zero cost.
// It removes the previous cost if this edge already exists.
func (g *PixelArtGraph) Add(v, w int) {
	g.AddCost(v, w, Color{})
}

// AddCost inserts a directed edge from v to w with cost c.
// It overwrites the previous cost if this edge already exists.
func (g *PixelArtGraph) AddCost(v, w int, c Color) {
	// Make sure not to break internal state.
	if w < 0 || w >= len(g.edges) {
		panic("vertex out of range: " + strconv.Itoa(w))
	}
	if g.edges[v] == nil {
		g.edges[v] = make(map[int]Color, initialMapSize)
	}
	g.edges[v][w] = c
}

// AddBoth inserts edges with zero cost between v and w.
// It removes the previous costs if these edges already exist.
func (g *PixelArtGraph) AddBoth(v, w int, c Color) {
	g.AddCost(v, w, c)
	if v != w {
		g.AddCost(w, v, c)
	}
}

// Delete removes an edge from v to w.
func (g *PixelArtGraph) Delete(v, w int) {
	delete(g.edges[v], w)
}

// DeleteBoth removes all edges between v and w.
func (g *PixelArtGraph) DeleteBoth(v, w int) {
	g.Delete(v, w)
	if v != w {
		g.Delete(w, v)
	}
}

// Degree returns the number of outward directed edges from v.
func (g *PixelArtGraph) Degree(v int) int {
	return len(g.edges[v])
}

// VisitColor calls the do function for each neighbor w of v,
// with c equal to the cost of the edge from v to w.
// The neighbors are visited in increasing numerical order.
// If do returns true, Visit returns immediately,
// skipping any remaining neighbors, and returns true.
func (g *PixelArtGraph) VisitColor(v int, do func(w int, c Color) bool) bool {
	for w, c := range g.edges[v] {
		if do(w, c) {
			return true
		}
	}
	return false
}

// Visit calls the do function for each neighbor w of v,
func (g *PixelArtGraph) Visit(v int, do func(w int) bool) bool {
	for w := range g.edges[v] {
		if do(w) {
			return true
		}
	}
	return false
}
