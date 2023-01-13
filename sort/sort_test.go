package sort

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBubble(t *testing.T) {
	tests := []struct {
		name     string
		sort     Sort
		input    []int
		expected []int
	}{
		{
			name:     "Bubble Sort",
			sort:     BubbleSort,
			input:    []int{52, 35, 0, 90, 99},
			expected: []int{0, 35, 52, 90, 99},
		},
		{
			name:     "Insert Sort",
			sort:     InsertSort,
			input:    []int{52, 35, 0, 90, 99},
			expected: []int{0, 35, 52, 90, 99},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.sort(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
