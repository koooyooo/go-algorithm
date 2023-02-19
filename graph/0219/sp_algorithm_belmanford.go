package _219

import "math"

func bellmanford(g map[int][]*edge, N, s int) (map[int]int, error) {
	// 距離を定義し初期化
	var dist = make(map[int]int)
	for i := 0; i < N; i++ {
		dist[i] = math.MaxInt32
	}
	// 始点は0距離として再設定
	dist[0] = 0

	for ite := 0; ite < N; ite++ {
		for i := 0; i < N; i++ {
			// 親の距離
			vDist := dist[i]
			// 親の距離が最大値であれば計算の意味がないので continue
			if vDist == math.MaxInt32 {
				continue
			}
			for _, e := range g[i] {
				// 子の距離
				toDist := dist[e.to]
				if w, ok := chmin(toDist, vDist+e.weight); ok {
					dist[e.to] = w
				}
			}
		}
	}

	// TODO 負閉路が無いかの判定が無い

	return dist, nil
}
