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
	graph := map[int]*point{
		1: {neighbors: []int{2, 3}},
		2: {neighbors: []int{1, 4, 5}},
		3: {neighbors: []int{1, 4}},
		4: {neighbors: []int{2, 3, 6}},
		5: {neighbors: []int{2, 6}},
		6: {neighbors: []int{4, 5}},
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
	graph := map[int]*point{
		1: {neighbors: []int{2, 3}},
		2: {neighbors: []int{1, 5}},
		3: {neighbors: []int{1, 4}},
		4: {neighbors: []int{3, 6}},
		5: {neighbors: []int{2}},
		6: {neighbors: []int{4}},
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
	graph := map[int]*point{
		1: {neighbors: []int{2, 3}},
		2: {neighbors: []int{1, 4, 5}},
		3: {neighbors: []int{1, 4}},
		4: {neighbors: []int{2, 3}},
		5: {neighbors: []int{2}},
		6: {neighbors: []int{}},
	}
	result := connected(graph, 1)
	assert.False(t, result)
}
