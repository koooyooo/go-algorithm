package _219

import (
	"fmt"
	"math"
)

// bellmanford はベルマンフォード法による最短経路の探索
func bellmanford(g map[int][]*edge, N, s int) (map[int]int, error) {

	// ① 距離を定義し初期化
	var dist = make(map[int]int)
	for i := 0; i < N; i++ {
		dist[i] = math.MaxInt32
	}
	dist[s] = 0

	// ② N回の基底ループを回しながら 更新を検知
	for ite := 0; ite < N; ite++ {
		var gUpdated bool

		// ③ N回の頂点ループを回しながら各頂点を最適化
		for i := 0; i < N; i++ {
			// ④ 親の距離が不明（最大値）の頂点からの経路は無視
			if dist[i] == math.MaxInt32 {
				continue
			}

			// ⑤ 子頂点の距離を 親頂点の距離+重みで緩和できれば A:距離を更新し B:更新を報告
			for _, e := range g[i] {
				if w, ok := chmin(dist[e.to], dist[i]+e.weight); ok {
					dist[e.to] = w
					gUpdated = true
				}
			}
		}

		// ⑥ 負閉路の検出（最短経路の辺の数は頂点の数より1つ少ないので、頂点の最大Idx[N-1]回目のサイクルで更新があれば 負閉路が存在）
		if ite == N-1 && gUpdated {
			return nil, fmt.Errorf("found negative cycle")
		}
	}

	return dist, nil
}
