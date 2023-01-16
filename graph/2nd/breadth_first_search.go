package main

import (
	"io"
)

func checkAllVisitedFromReaderBFS(r io.Reader) bool {
	lines := read(r)
	graph := toGraph(lines[1:])
	return checkAllVisitedBFS(graph)
}

func checkAllVisitedBFS(graph map[int]*point) bool {
	var queue []*point
	queue = append(queue, graph[1])
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		head.visit = true
		for _, n := range head.neighbors {
			point := graph[n]
			if point.visit {
				continue
			}
			queue = append(queue, point)
		}
	}
	for _, v := range graph {
		if !v.visit {
			return false
		}
	}
	return true
}
