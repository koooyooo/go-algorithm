package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	tests := []struct {
		name  string
		input string
		exp   bool
	}{
		{
			name: "頂点数: 2, 辺の数: 1, 綿棒: 二郎グラフ",
			input: `2 1
1 2`,
			exp: true,
		}, {
			name: "頂点数: 4, 辺の数: 4, 四角: 二郎グラフ",
			input: `4 4
1 2
2 3
3 4
4 1`,
			exp: true,
		}, {
			name: "頂点数: 4, 辺の数: 6, 四角 + 対角線: 非二郎グラフ",
			input: `4 6
1 2
2 3
3 4
4 1
1 3
2 4`,
			exp: false,
		}, {
			name: "頂点数: 3, 辺の数: 3, 三角: 非二郎グラフ",
			input: `3 3
1 2
2 3
3 1`,
			exp: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.input)
			lines := readInput(r)
			graph := createGraph(lines[1:])
			result := checkNon2Color(graph)
			assert.Equal(t, test.exp, result)
		})
	}
}
