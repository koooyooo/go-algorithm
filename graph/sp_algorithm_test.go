package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	g := map[int][]*edge{
		0: {{to: 1, weight: 3}, {to: 2, weight: 5}},
		1: {{to: 2, weight: 4}, {to: 3, weight: 12}},
		2: {{to: 3, weight: 9}, {to: 4, weight: 4}},
		3: {{to: 5, weight: 2}},
		4: {{to: 3, weight: 7}, {to: 5, weight: 8}},
		5: nil,
	}

	tests := []struct {
		name string
		alg  shortestPathProblem
	}{
		{
			name: "bellmanford",
			alg:  bellmanford,
		},
		{
			name: "dijkstrasLoop",
			alg:  dijkstrasLoop,
		},
		{
			name: "dijkstrasQueue",
			alg:  dijkstrasQueue,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			alg := test.alg
			d, err := alg(g, len(g), 0)
			assert.NoError(t, err)

			assert.Equal(t, 0, d[0])
			assert.Equal(t, 3, d[1])
			assert.Equal(t, 5, d[2])
			assert.Equal(t, 14, d[3])
			assert.Equal(t, 9, d[4])
			assert.Equal(t, 16, d[5])
		})
	}

}
