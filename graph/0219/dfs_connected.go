package _219

// グラフの連結を確認する
func connected(g map[int][]int, s int) bool {
	// ① 訪問履歴 Mapを初期化
	visited := make(map[int]bool)
	// ② スタート地点から再帰的に走査
	visit(g, visited, s)
	// ③ 結果確認（訪問履歴の無い頂点が残っているか）
	for v, _ := range g {
		if !visited[v] {
			return false
		}
	}
	return true
}

// 目標がリストに無ければ一段回戻るロジックは再起で表現できる
func visit(g map[int][]int, visited map[int]bool, current int) {
	// ④ 訪問履歴 を記載
	visited[current] = true
	// ⑤ 周囲の頂点を再帰的に検索
	for _, p := range g[current] {
		// ⑥ 訪問履歴があればスキップ
		if visited[p] {
			continue
		}
		visit(g, visited, p)
	}
}
