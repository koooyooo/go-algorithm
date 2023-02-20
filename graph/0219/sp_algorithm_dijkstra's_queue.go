package _219

import "math"

// elm はQueueの要素
type elm struct {
	v     int
	vDist int
}

// popNearest 最短距離の要素をPopする関数
func popNearest(pq []*elm) (*elm, []*elm) {
	if len(pq) == 0 {
		return nil, pq
	}
	var shortestIdx = -1
	var shortest *elm
	for i, v := range pq {
		if shortest == nil || shortest.vDist > v.vDist {
			shortestIdx = i
			shortest = v
		}
	}
	return shortest, append(pq[0:shortestIdx], pq[shortestIdx+1:]...)
}

func dijkstrasQueue(g map[int][]*edge, N, s int) (map[int]int, error) {
	// 距離の初期化
	var dist = make(map[int]int, N)
	for i := 0; i < N; i++ {
		dist[i] = math.MaxInt32
	}
	dist[s] = 0

	// Queueの初期化
	var queue = []*elm{{v: s, vDist: 0}}
	for len(queue) != 0 {
		// 最短経路の要点
		var nearestElm *elm
		nearestElm, queue = popNearest(queue)

		for _, e := range g[nearestElm.v] {
			// 各辺を辿った経路で最短値が更新できたら、Queue追加と距離更新を行う
			if min, ok := chmin(dist[e.to], dist[nearestElm.v]+e.weight); ok {
				queue = append(queue, &elm{
					v:     e.to,
					vDist: min,
				})
				dist[e.to] = min
			}
		}
	}
	return dist, nil
}
