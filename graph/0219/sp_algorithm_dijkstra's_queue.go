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
	// 最短距離を探し出し
	var shortestIdx = -1
	var shortest *elm
	for i, v := range pq {
		if shortest == nil || shortest.vDist > v.vDist {
			shortestIdx = i
			shortest = v
		}
	}
	// 最短距離と最短距離を除いた残QueueをReturn
	return shortest, append(pq[0:shortestIdx], pq[shortestIdx+1:]...)
}

func dijkstrasQueue(g map[int][]*edge, N, s int) (map[int]int, error) {

	// ① 距離の初期化
	var dist = make(map[int]int, N)
	for i := 0; i < N; i++ {
		dist[i] = math.MaxInt32
	}
	dist[s] = 0

	// ② 最短距離優先 Queueの初期化
	var queue = []*elm{{v: s, vDist: 0}}

	for len(queue) != 0 {
		// ③ 最短距離の頂点をPOP
		var nearestElm *elm
		nearestElm, queue = popNearest(queue)

		// ④ 各辺を辿って緩和
		for _, e := range g[nearestElm.v] {
			if min, ok := chmin(dist[e.to], dist[nearestElm.v]+e.weight); ok {
				// ⑤ 緩和時の処理は A:Queue追加 と B:距離更新
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
