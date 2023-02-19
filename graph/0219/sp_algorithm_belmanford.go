package _219

import (
	"fmt"
	"math"
)

func bellmanford(g map[int][]*edge, N, s int) (map[int]int, error) {
	// 距離を定義し初期化
	var dist = make(map[int]int)
	for i := 0; i < N; i++ {
		dist[i] = math.MaxInt32
	}
	// 始点は0距離として再設定
	dist[s] = 0

	for ite := 0; ite < N; ite++ {
		var gUpdated bool

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
					gUpdated = true
				}
			}
		}

		// 負閉路の検出（最短経路の辺の数は頂点の数より1つ少ないので、頂点の最大Idx[N-1]回目のサイクルで更新があれば 負閉路が存在）
		if ite == N-1 && gUpdated {
			return nil, fmt.Errorf("found negative cycle")
		}
	}

	return dist, nil
}
