package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type point struct {
	visited  bool
	white    bool
	neighbor []int
}

func main() {
	lines := readInput(os.Stdin)
	graph := createGraph(lines[1:])
	result := checkNon2Color(graph)
	fmt.Printf("2 color : %t", result)
}

func createGraph(lines []string) map[int]*point {
	graph := make(map[int]*point)
	for _, l := range lines {
		elms := strings.Split(l, " ")
		ai, _ := strconv.Atoi(elms[0])
		bi, _ := strconv.Atoi(elms[1])
		if _, ok := graph[ai]; !ok {
			graph[ai] = &point{}
		}
		graph[ai].neighbor = append(graph[ai].neighbor, bi)
		if _, ok := graph[bi]; !ok {
			graph[bi] = &point{}
		}
		graph[bi].neighbor = append(graph[bi].neighbor, ai)
	}
	return graph
}

func checkNon2Color(graph map[int]*point) bool {
	var queue []*point
	queue = append(queue, graph[1])

	for len(queue) > 0 {
		head := queue[0]
		head.visited = true
		queue = queue[1:]

		for _, n := range head.neighbor {
			neighbor := graph[n]
			if neighbor.visited {
				if neighbor.white == head.white {
					return false
				}
				continue
			}
			neighbor.white = !head.white
			queue = append(queue, neighbor)
		}
	}
	return true
}

func readInput(r io.Reader) []string {
	var lines []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}
