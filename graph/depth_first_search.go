package graph

type point struct {
	visited   bool
	neighbors []int
}

// グラフの連結を確認する
func connected(graph map[int]*point, init int) bool {
	initPoint := graph[init]
	// スタート地点から枝葉を再帰的に走査しvisitedマークを付けていく
	visit(graph, initPoint)
	// 走査後にvisitedマークの無い箇所が無ければグラフは連結されている
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
