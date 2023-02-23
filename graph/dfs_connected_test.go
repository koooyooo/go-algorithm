package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// [1]から開始しグラフが連結していることを確認する
//
// [1]--[2]--[5]
//
//	|    |    |
//
// [3]--[4]--[6]
func TestDepthFirst(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 4, 5},
		3: {1, 4},
		4: {2, 3, 6},
		5: {2, 6},
		6: {4, 5},
	}
	result := connected(graph, 1)
	assert.True(t, result)
}

// [1]から開始しグラフが連結していることを確認する
//
// [1]--[2]--[5]
//
//	|
//
// [3]--[4]--[6]
func TestDepthFirstConnected(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 5},
		3: {1, 4},
		4: {3, 6},
		5: {2},
		6: {4},
	}
	result := connected(graph, 1)
	assert.True(t, result)
}

// [1]から開始しグラフが連結されていないことを確認する
//
// [1]--[2]--[5]
//
//	|    |
//
// [3]--[4]  [6]
func TestDepthFirstNotConnected(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 4, 5},
		3: {1, 4},
		4: {2, 3},
		5: {2},
		6: {},
	}
	result := connected(graph, 1)
	assert.False(t, result)
}
