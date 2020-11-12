package main

import (
	"strconv"
)

const initialMapSize = 4

// Graph with pixel information
type Graph struct {
	edges []map[int]Color
}

// NewGraph constructs a new graph with n vertices, numbered from 0 to n-1, and no edges.
func NewGraph(n int) *Graph {
	return &Graph{edges: make([]map[int]Color, n)}
}

// Order returns the number of vertices in the graph.
func (g *Graph) Order() int {
	return len(g.edges)
}

// Edge tells if there is an edge from v to w.
func (g *Graph) Edge(v, w int) bool {
	if v < 0 || v >= g.Order() {
		return false
	}
	_, ok := g.edges[v][w]
	return ok
}

// Add inserts a directed edge from v to w with zero cost.
// It removes the previous cost if this edge already exists.
func (g *Graph) Add(v, w int) {
	g.AddCost(v, w, Color{})
}

// AddCost inserts a directed edge from v to w with cost c.
// It overwrites the previous cost if this edge already exists.
func (g *Graph) AddCost(v, w int, c Color) {
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
func (g *Graph) AddBoth(v, w int) {
	g.Add(v, w)
	if v != w {
		g.Add(w, v)
	}
}

// Delete removes an edge from v to w.
func (g *Graph) Delete(v, w int) {
	delete(g.edges[v], w)
}

// DeleteBoth removes all edges between v and w.
func (g *Graph) DeleteBoth(v, w int) {
	g.Delete(v, w)
	if v != w {
		g.Delete(w, v)
	}
}

// Degree returns the number of outward directed edges from v.
func (g *Graph) Degree(v int) int {
	return len(g.edges[v])
}

// VisitColor calls the do function for each neighbor w of v,
// with c equal to the cost of the edge from v to w.
// The neighbors are visited in increasing numerical order.
// If do returns true, Visit returns immediately,
// skipping any remaining neighbors, and returns true.
func (g *Graph) VisitColor(v int, do func(w int, c Color) bool) bool {
	for w, c := range g.edges[v] {
		if do(w, c) {
			return true
		}
	}
	return false
}

// Visit calls the do function for each neighbor w of v,
func (g *Graph) Visit(v int, do func(w int) bool) bool {
	for w := range g.edges[v] {
		if do(w) {
			return true
		}
	}
	return false
}
