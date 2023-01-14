package graph

type distancePoint struct {
	distance int
	point
}

func fillDistance(graph map[int]*distancePoint, init int) {
	initPoint := graph[init]
	initPoint.distance = 0
	queue := []*distancePoint{initPoint}
	handleQueue(graph, queue)
}

func handleQueue(graph map[int]*distancePoint, queue []*distancePoint) {
	for len(queue) > 0 {
		head := queue[0]
		head.visited = true
		queue = queue[1:] // Queueは計算対象なので、自身は除外しておく
		for _, idx := range head.neighbors {
			neighbor := graph[idx]
			if neighbor.visited { // 処理済のPointは追加自体をしない（追加後に処理をしないロジックだと 空の仕事がQueueにたまり続けて無限ループになる）
				continue
			}
			neighbor.distance = head.distance + 1 // 親の値が見えるこの時点でインクリメント
			queue = append(queue, neighbor)
		}
	}
}
