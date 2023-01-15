package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount1SmallerLinkPoint(t *testing.T) {
	// 4, 5 が対象
	lines := []line{
		{0, 2},
		{1, 2},
		{1, 3},
		{2, 3},
		{2, 4},
		{2, 5},
	}
	result := count1SmallerLinkPoint(lines)
	assert.Equal(t, 2, result)
}
