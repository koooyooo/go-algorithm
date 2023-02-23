package _st

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
	graph := map[int]*distancePoint{
		1: {point: point{
			neighbors: []int{2, 3},
		}},
		2: {point: point{
			neighbors: []int{1, 4, 5},
		}},
		3: {point: point{
			neighbors: []int{1, 4},
		}},
		4: {point: point{
			neighbors: []int{2, 3, 6},
		}},
		5: {point: point{
			neighbors: []int{2, 6},
		}},
		6: {point: point{
			neighbors: []int{4, 5},
		}},
	}
	fillDistance(graph, 1)
	assert.Equal(t, graph[1].distance, 0)
	assert.Equal(t, graph[2].distance, 1)
	assert.Equal(t, graph[3].distance, 1)
	assert.Equal(t, graph[4].distance, 2)
	assert.Equal(t, graph[5].distance, 2)
	assert.Equal(t, graph[6].distance, 3)
}
