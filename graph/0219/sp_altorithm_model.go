package _219

// map[int]*v --- []int

type edge struct {
	to     int
	weight int
}

type shortestPathProblem func(g map[int][]*edge, N int, s int) (map[int]int, error)

func chmin(a, b int) (int, bool) {
	if a > b {
		return b, true
	}
	return a, false
}
