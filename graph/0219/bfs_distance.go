package _219

// [幅優先探索]
// スタート地点からの距離を測定
func fillDistance(g map[int][]int, s int) map[int]int {
	// 距離と発見済みマークを初期化
	dist := map[int]int{s: 0}
	seen := map[int]bool{s: true}

	queue := []int{s}

	for len(queue) > 0 {
		// 先頭要素をPOP
		head := queue[0]
		queue = queue[1:]

		for _, n := range g[head] {
			// 発見済の場合はキャンセル（Queue追加による無限ループ予防）
			if seen[n] {
				continue
			}
			// 実処理
			dist[n] = dist[head] + 1 // 親の値が見えるこの時点でインクリメント

			// 発見済みとしてマークした上で Queueに追加（①未発見 ②発見済-Queue処理待ち ③発見済-Queue処理済 …のステップにおける②）
			seen[n] = true
			queue = append(queue, n)
		}
	}
	return dist
}
