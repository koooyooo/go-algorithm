package graph

import (
	"fmt"
	"math"
)

type E struct {
	To     int
	Weight int
}

func (e *E) String() string {
	return fmt.Sprintf("--(%d)-->%d", e.Weight, e.To)
}

type QE struct {
	V        int
	Distance int
}

type Graph map[int][]*E

func chmin(a int, b int) (bool, int) {
	if a > b {
		return true, b
	}
	return false, a
}

func FillDistance(g Graph, s int) map[int]int {
	// 距離マップを初期化
	var dist = make(map[int]int)
	for v, _ := range g {
		init := math.MaxInt32
		if v == s {
			init = 0
		}
		dist[v] = init
	}

	// 最短経路を初期化
	queue := []*QE{{s, 0}}
	for len(queue) != 0 {
		route := queue[0]
		queue = queue[1:]

		// 次の頂点の重みを辺の重みだけで上回った場合はゴミルート
		if dist[route.V] < route.Distance {
			continue
		}

		for _, e := range g[route.V] {
			// 最短経路を更新した頂点に関しては作業キューに入れる
			if updated, min := chmin(dist[e.To], dist[route.V]+e.Weight); updated {
				queue = append(queue, &QE{
					V:        e.To,
					Distance: min,
				})
				dist[e.To] = min
			}
		}
	}
	return dist
}
