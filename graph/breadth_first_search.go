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
		queue = queue[1:]
		for _, idx := range head.neighbors {
			neighbor := graph[idx]
			if neighbor.visited {
				continue
			}
			neighbor.distance = head.distance + 1
			queue = append(queue, neighbor)
		}
	}
}
