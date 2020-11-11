package main

import (
	"strconv"
)

// Graph with pixel information
type Graph struct {
	edges [][]int
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
	if w < 0 || w >= len(g.edges[v]) {
		return false
	}
	ok := g.edges[v][w] > 0
	return ok
}

// Add inserts a directed edge from v to w with cost c.
// It overwrites the previous cost if this edge already exists.
func (g *Graph) Add(v, w int) {
	// Make sure not to break internal state.
	if w < 0 || w >= len(g.edges) {
		panic("vertex out of range: " + strconv.Itoa(w))
	}
	if g.edges[v] == nil {
		g.edges[v] = make([]int, 0)
	}
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
	// Remove the element at index i from a.
	g.edges[v][w] = g.edges[v][len(g.edges[v])-1] // Copy last element to index i.
	g.edges[v][len(g.edges[v])-1] = 0             // Erase last element (write zero value).
	g.edges[v] = g.edges[v][:len(g.edges[v])-1]   // Truncate slice.
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

// Visit calls the do function for each neighbor w of v,
// with c equal to the cost of the edge from v to w.
// The neighbors are visited in increasing numerical order.
// If do returns true, Visit returns immediately,
// skipping any remaining neighbors, and returns true.
func (g *Graph) Visit(v int, do func(w int) bool) bool {
	for _, e := range g.edges[v] {
		if do(e) {
			return true
		}
	}
	return false
}
