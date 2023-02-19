package _219

import "math"

func dijkstrasN2(g map[int][]*edge, N, s int) (map[int]int, error) {
	// 最短頂点利用済みフラグ
	seen := make(map[int]bool)
	// 距離の初期化
	dist := make(map[int]int)
	for i := 0; i < N; i++ {
		dist[i] = math.MaxInt32
	}
	dist[s] = 0

	for iter := 0; iter < N; iter++ {
		minDistance := math.MaxInt32
		minV := -1

		// 最小値を更新
		for v := 0; v < N; v++ {
			// 未確認で最小の距離を持つ頂点を漁る
			if !seen[v] && dist[v] < minDistance {
				minV = v
				minDistance = dist[minV]
			}
		}

		// 未確認の中に最小値が見つからなければクリア
		noMinFound := minV == -1
		if noMinFound {
			break
		}

		// 子の頂点にそれぞれ値を設定
		for _, e := range g[minV] {
			// 子の最小値を更新できそうなら更新
			if dist[e.to] > dist[minV]+e.weight {
				dist[e.to] = dist[minV] + e.weight
			}
		}
		// 処理済マークに更新
		seen[minV] = true
	}

	return dist, nil
}
