package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader(`o - o o o o
o - o o - o
o - o o - o
o o o o o o
- - - - o o
o o o o o o`)
	lines := readInput(input)
	matrix := createBoolMatrix(lines)
	graph := createGraph(matrix)
	fillDistance(graph)
	fmt.Printf("endpoint distance: %d\n", graph[address{5, 5}].distance)
}

func fillDistance(g map[address]*point) {
	var queue []*point
	queue = append(queue, g[address{0, 0}])
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		head.visit = true

		for _, n := range head.neighbors {
			if n.visit {
				continue
			}
			n.distance = head.distance + 1
			queue = append(queue, n)
		}
	}
}

type address struct {
	row int
	col int
}

type point struct {
	address   address
	passable  bool
	visit     bool
	distance  int
	neighbors []*point
}

func createGraph(m [][]bool) map[address]*point {
	minRow := 0
	maxRow := len(m) - 1
	minCol := 0
	maxCol := len(m[0]) - 1
	graph := map[address]*point{}
	for i, r := range m {
		for j, c := range r {
			a := address{i, j}
			graph[a] = &point{
				address:  a,
				passable: c,
			}
		}
	}
	var inMap = func(row, col int) bool {
		if row < minRow || maxRow < row {
			return false
		}
		if col < minCol || maxCol < col {
			return false
		}
		return true
	}
	var addNeighborIfOK = func(basePoint *point, row, col int) {
		if inMap(row, col) {
			point := graph[address{row, col}]
			if point.passable {
				basePoint.neighbors = append(basePoint.neighbors, point)
			}
		}
	}
	for a, p := range graph {
		addNeighborIfOK(p, a.row-1, a.col)
		addNeighborIfOK(p, a.row+1, a.col)
		addNeighborIfOK(p, a.row, a.col-1)
		addNeighborIfOK(p, a.row, a.col+1)
	}
	return graph
}

func createBoolMatrix(line []string) [][]bool {
	var mLine [][]bool
	for _, l := range line {
		cols := strings.Split(l, " ")
		var mCol []bool
		for _, c := range cols {
			v := c == "o"
			mCol = append(mCol, v)
		}
		mLine = append(mLine, mCol)
	}
	return mLine
}

func readInput(r io.Reader) (lines []string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return
}
