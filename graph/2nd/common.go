package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type point struct {
	visit     bool
	neighbors []int
}

// [1]--[2]--[5]
//
//	|    |    |
//
// [3]--[4]--[6]
//
// [Input]
// M:頂点数 N:辺数
// 頂点A 頂点B
var input = `6 7
1 2
2 5
1 3
2 4
5 6
3 4
4 6`

func toGraph(lines []string) map[int]*point {
	graph := make(map[int]*point)
	for _, l := range lines {
		elms := strings.Split(l, " ")
		a, _ := strconv.Atoi(elms[0])
		b, _ := strconv.Atoi(elms[1])

		pointA, ok := graph[a]
		if !ok {
			pointA = &point{}
			graph[a] = pointA
		}
		pointA.neighbors = append(pointA.neighbors, b)
		pointB, ok := graph[b]
		if !ok {
			pointB = &point{}
			graph[b] = pointB
		}
		pointB.neighbors = append(pointB.neighbors, a)
	}
	return graph
}

func read(r io.Reader) []string {
	var lines []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}
