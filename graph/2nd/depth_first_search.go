package main

import (
	"io"
)

func checkAllVisitedFromReaderDFS(r io.Reader) bool {
	lines := read(r)
	graph := toGraph(lines[1:])
	return checkAllVisitedDFS(graph)
}

func checkAllVisitedDFS(graph map[int]*point) bool {
	visit(graph, 1)
	for _, p := range graph {
		if !p.visit {
			return false
		}
	}
	return true
}

func visit(g map[int]*point, start int) {
	point := g[start]
	point.visit = true
	for _, n := range point.neighbors {
		if g[n].visit {
			continue
		}
		visit(g, n)
	}
}
