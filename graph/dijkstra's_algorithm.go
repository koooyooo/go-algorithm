package graph

import (
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

func (e *Edge) String() string {
	return fmt.Sprintf("--(%d)-->%d", e.Weight, e.To)
}

type RouteV struct {
	V        int
	Distance int
}

type Graph map[int][]*Edge

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

	// 検索候補キューを初期化
	queue := []*RouteV{{s, 0}}
	for len(queue) != 0 {
		route := queue[0]
		queue = queue[1:]

		// 次の頂点の重みを辺の重み単体で上回った場合はゴミルート
		if dist[route.V] < route.Distance {
			continue
		}

		for _, e := range g[route.V] {
			// 最短距離を更新した頂点に関しては
			// - 距離を更新し
			// - 検索候補キューに入れる
			if updated, min := chmin(dist[e.To], dist[route.V]+e.Weight); updated {
				dist[e.To] = min
				queue = append(queue, &RouteV{
					V:        e.To,
					Distance: min,
				})
			}
		}
	}
	return dist
}
