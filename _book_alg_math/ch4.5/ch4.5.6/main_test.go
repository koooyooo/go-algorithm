package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFillDistanceSingleRoute(t *testing.T) {
	input := strings.NewReader(`o - o
o - o
o o o`)
	lines := readInput(input)
	matrix := createBoolMatrix(lines)
	graph := createGraph(matrix)
	fillDistance(graph)
	assert.Equal(t, 0, graph[address{0, 0}].distance)
	assert.Equal(t, 1, graph[address{1, 0}].distance)
	assert.Equal(t, 2, graph[address{2, 0}].distance)
	assert.Equal(t, 3, graph[address{2, 1}].distance)
	assert.Equal(t, 4, graph[address{2, 2}].distance)
	assert.Equal(t, 5, graph[address{1, 2}].distance)
	assert.Equal(t, 6, graph[address{0, 2}].distance)
}

func TestFillDistanceDoubleRoute(t *testing.T) {
	input := strings.NewReader(`o o o
o - o
o o o`)
	lines := readInput(input)
	matrix := createBoolMatrix(lines)
	graph := createGraph(matrix)
	fillDistance(graph)
	assert.Equal(t, 0, graph[address{0, 0}].distance)
	assert.Equal(t, 1, graph[address{1, 0}].distance)
	assert.Equal(t, 1, graph[address{0, 1}].distance)
	assert.Equal(t, 2, graph[address{2, 0}].distance)
	assert.Equal(t, 2, graph[address{0, 2}].distance)
	assert.Equal(t, 3, graph[address{2, 1}].distance)
	assert.Equal(t, 3, graph[address{1, 2}].distance)
	assert.Equal(t, 4, graph[address{2, 2}].distance)
}

func TestFillDistanceAllRoute(t *testing.T) {
	input := strings.NewReader(`o o o
o o o
o o o`)
	lines := readInput(input)
	matrix := createBoolMatrix(lines)
	graph := createGraph(matrix)
	fillDistance(graph)
	assert.Equal(t, 0, graph[address{0, 0}].distance)
	assert.Equal(t, 1, graph[address{1, 0}].distance)
	assert.Equal(t, 1, graph[address{0, 1}].distance)
	assert.Equal(t, 2, graph[address{1, 1}].distance)
	assert.Equal(t, 2, graph[address{2, 0}].distance)
	assert.Equal(t, 2, graph[address{0, 2}].distance)
	assert.Equal(t, 3, graph[address{2, 1}].distance)
	assert.Equal(t, 3, graph[address{1, 2}].distance)
	assert.Equal(t, 4, graph[address{2, 2}].distance)
}
