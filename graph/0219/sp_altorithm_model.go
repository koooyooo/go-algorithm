package _219

// shortestPathProblem は最短経路探索のアルゴリズム
type shortestPathProblem func(g map[int][]*edge, N int, s int) (map[int]int, error)

// edge は重み付きの辺
type edge struct {
	to     int
	weight int
}

// chmin は既定値aと 更新値bを受け取り、bが最短だった場合は最短値を更新する
func chmin(a, b int) (int, bool) {
	if a > b {
		return b, true
	}
	return a, false
}
