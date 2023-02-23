package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// [1]から開始し全てのグラフの距離を算出する
//
// [1]--[2]--[5]
//
//	|    |    |
//
// [3]--[4]--[6]
func TestFillDistance(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 4, 5},
		3: {1, 4},
		4: {2, 3, 6},
		5: {2, 6},
		6: {4, 5},
	}
	dist := fillDistance(graph, 1)
	assert.Equal(t, 0, dist[1])
	assert.Equal(t, 1, dist[2])
	assert.Equal(t, 1, dist[3])
	assert.Equal(t, 2, dist[4])
	assert.Equal(t, 2, dist[5])
	assert.Equal(t, 3, dist[6])
}
