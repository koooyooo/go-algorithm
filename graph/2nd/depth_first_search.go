package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader(input)
	lines := read(r)
	graph := toGraph(lines[1:])
	visit(graph, 1)
	for _, p := range graph {
		if !p.visit {
			fmt.Println("Not All Visited")
			return
		}
	}
	fmt.Println("All Visited")
}

// mapはコピーではなく参照となるか
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
