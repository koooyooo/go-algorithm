package _st

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	g := Graph{
		0: []*Edge{{To: 1, Weight: 3}, {To: 2, Weight: 5}},
		1: []*Edge{{To: 2, Weight: 4}, {To: 3, Weight: 12}},
		2: []*Edge{{To: 3, Weight: 9}, {To: 4, Weight: 4}},
		3: []*Edge{{To: 5, Weight: 2}},
		4: []*Edge{{To: 3, Weight: 7}, {To: 5, Weight: 8}},
		5: nil,
	}
	vDist := FillDistance(g, 0)
	for v, i := range vDist {
		fmt.Printf("%d: %d\n", v, i)
	}
}
