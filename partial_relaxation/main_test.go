package partial_relaxation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelax(t *testing.T) {
	tests := []struct {
		name string
		i    int
		w    int
		a    []int
		exp  bool
	}{
		{
			name: "[3, 6, 5] の3点を用いれば 14を作れる",
			i:    4,
			w:    14,
			a:    []int{3, 2, 6, 5},
			exp:  true,
		},
		{
			name: "[3, 6, 5] の2点では 14を作れない",
			i:    2,
			w:    14,
			a:    []int{3, 2, 6, 5},
			exp:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := relax(test.i, test.w, test.a)
			assert.Equal(t, test.exp, res)
		})
	}

}
