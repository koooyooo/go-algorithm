package _219

type pathV struct {
	v     int
	vDist int
}

type pathQueue []*pathV

func popShortest(pq pathQueue) (*pathV, pathQueue) {
	if len(pq) == 0 {
		return nil, pq
	}
	var shortest *pathV
	for _, v := range pq {
		if shortest == nil || shortest.vDist > v.vDist {
			shortest = v
		}
	}
	return shortest, pq[1:]
}

func dijkstrasQueue(g map[int][]*edge, N, s int) (map[int]int, error) {

	return nil, nil
}
