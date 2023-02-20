package _219

import "math"

// dijkstrasLoop はダイクストラ法による距離測定
func dijkstrasLoop(g map[int][]*edge, N, s int) (map[int]int, error) {
	// ① 距離の初期化
	dist := make(map[int]int)
	for i := 0; i < N; i++ {
		dist[i] = math.MaxInt32
	}
	dist[s] = 0

	// ② 最短頂点処理済みMapを用意
	seen := make(map[int]bool)

	// ③ N回の基底ループを回す
	for iter := 0; iter < N; iter++ {

		// ④ 最短頂点とその距離を N回のループで検索
		minV := -1
		minDistance := math.MaxInt32

		for v := 0; v < N; v++ {
			// 未確認 & 最短距離の頂点が対象
			if !seen[v] && dist[v] < minDistance {
				minV = v
				minDistance = dist[minV]
			}
		}

		// ⑤ 更新対象が無ければ全て処理済み
		if minV == -1 {
			break
		}

		// ⑥ 最短頂点から伸びる辺の先を緩和
		for _, e := range g[minV] {
			if min, ok := chmin(dist[e.to], dist[minV]+e.weight); ok {
				dist[e.to] = min
			}
		}
		// ⑦ 最短頂点を処理済としてマーク
		seen[minV] = true
	}
	return dist, nil
}
