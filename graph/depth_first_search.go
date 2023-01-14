package graph

type point struct {
	visited   bool
	neighbors []int
}

func connected(graph map[int]*point, init int) bool {
	initPoint := graph[init]
	visit(graph, initPoint)
	for _, p := range graph {
		if p.visited == false {
			return false
		}
	}
	return true
}

// 目標がリストに無ければ一段回戻るロジックは再起で表現できる
func visit(graph map[int]*point, currentPoint *point) {
	currentPoint.visited = true
	for _, p := range currentPoint.neighbors {
		if !graph[p].visited {
			visit(graph, graph[p])
		}
	}
}
