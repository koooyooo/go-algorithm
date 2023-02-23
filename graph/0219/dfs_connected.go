package _219

// グラフの連結を確認する
func connected(g map[int][]int, s int) bool {
	visited := make(map[int]bool)
	// スタート地点から走査する
	visit(g, visited, s)
	// visitedフラグの無い頂点があれば未到達ありと判断
	for v, _ := range g {
		if !visited[v] {
			return false
		}
	}
	return true
}

// 目標がリストに無ければ一段回戻るロジックは再起で表現できる
func visit(g map[int][]int, visited map[int]bool, current int) {
	visited[current] = true
	for _, p := range g[current] {
		if visited[p] {
			continue
		}
		visit(g, visited, p)
	}
}
