package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var inputStr = `5 6
0 2
1 2
1 3
2 3
2 4
2 5
`

func main() {
	in := strings.NewReader(inputStr)
	// 頂点の数"M"は後に生成するMapのキーから、辺の数"N"はlineの行数で判明するので取得不要
	raw := readInput(in)
	lines := toLine(raw[1:])
	count := count1SmallerLinkPoint(lines)
	fmt.Println(count)
}

type line struct {
	a int
	b int
}

func count1SmallerLinkPoint(lines []line) int {
	countMap := map[int]int{}
	for _, l := range lines {
		if l.a > l.b {
			countMap[l.a]++
		}
		if l.a < l.b {
			countMap[l.b]++
		}
	}
	pointCount := 0
	for _, v := range countMap {
		if v == 1 {
			pointCount++
		}
	}
	return pointCount
}

// [][]string の構造体は中間的すぎるので排除、モデルを適用
func toLine(lines []string) []line {
	var ls []line
	for _, l := range lines {
		cols := strings.Split(l, " ")
		a, _ := strconv.Atoi(cols[0])
		b, _ := strconv.Atoi(cols[1])
		ls = append(ls, line{
			a: a,
			b: b,
		})
	}
	return ls
}

func readInput(r io.Reader) []string {
	s := bufio.NewScanner(r)
	var line []string
	for s.Scan() {
		line = append(line, s.Text())
	}
	return line
}
