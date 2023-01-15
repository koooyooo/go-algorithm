package graph

type distancePoint struct {
	distance int
	point
}

// [幅優先探索]
// スタート地点からの距離を測定
func fillDistance(graph map[int]*distancePoint, init int) {
	initPoint := graph[init]
	initPoint.distance = 0
	queue := []*distancePoint{initPoint}
	handleQueue(graph, queue)
}

func handleQueue(graph map[int]*distancePoint, queue []*distancePoint) {
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]   // 計算済みとなる自身はQueueから除外
		head.visited = true // 訪問済みマークを忘れずに
		for _, idx := range head.neighbors {
			neighbor := graph[idx]
			if neighbor.visited { // 訪問済みマークがあれば詰め込まない（Queueに詰め込んだ後で対応しようとしても時すでに遅しで無限追加に入る）
				continue
			}
			neighbor.distance = head.distance + 1 // 親の値が見えるこの時点でインクリメント
			queue = append(queue, neighbor)
		}
	}
}
