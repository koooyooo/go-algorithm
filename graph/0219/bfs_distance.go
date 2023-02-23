package _219

// [幅優先探索]
// スタート地点からの距離を測定
func fillDistance(g map[int][]int, s int) map[int]int {
	// ① 拡張点の距離・発見済Mapを初期化・開始地点を設定
	dist := map[int]int{s: 0}
	seen := map[int]bool{s: true}

	// ② Queueを初期化・開始地点を設定
	queue := []int{s}

	// ③ Queueの要素でループ（終了条件 = Queue要素なし）
	for len(queue) > 0 {

		// ④ 先頭要素をPOP
		head := queue[0]
		queue = queue[1:]

		// ⑤ 先頭要素の周囲を探索
		for _, n := range g[head] {
			// ⑥ 発見済の頂点はキャンセル（Queue追加による無限ループ予防）
			if seen[n] {
				continue
			}
			// ⑦ 実処理（距離測定）
			dist[n] = dist[head] + 1 // 親の値が見えるこの時点でインクリメント

			// ⑧ 頂点を発見済みとしてマークして Queueに追加
			//  （①未発見 ②発見済-Queue処理待ち ③発見済-Queue処理済 …のステップにおける②）
			seen[n] = true
			queue = append(queue, n)
		}
	}
	return dist
}
