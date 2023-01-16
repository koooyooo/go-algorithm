package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// [1]--[2]--[5]
//
//	|    |    |
//
// [3]--[4]--[6]
//
// [Input]
// M:頂点数 N:辺数
// 頂点A 頂点B
func TestCheckAllVisitedFromReaderDFS(t *testing.T) {
	tests := []struct {
		name  string
		input string
		exp   bool
	}{
		{
			name: "all visited",
			input: `6 7
1 2
2 5
1 3
2 4
5 6
3 4
4 6`,
			exp: true,
		},
		{
			name: "not all visited",
			input: `7 7
1 2
2 5
1 3
2 4
5 6
3 4
4 6
7 7`,
			exp: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkAllVisitedFromReaderDFS(strings.NewReader(test.input))
			assert.Equal(t, test.exp, result)
		})
	}

}
