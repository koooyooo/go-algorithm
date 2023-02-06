package _basic_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	type person struct {
		friends []int
	}
	p := person{}
	fs := p.friends
	for i := 0; i < 10; i++ {
		fs = append(fs, i)
	}
	// 属性値のコピーを操作してもコピーは変更されるが
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, fs)
	// structは変更されない
	assert.Equal(t, []int{}, p.friends)
}

// 関数に引数として渡したSliceを関数内で操作しても
// 関数外には反映されない
func TestSliceArgUpdate(t *testing.T) {
	f := func(s []int) {
		s = append(s, 1)
	}
	in := []int{}
	f(in)
	assert.Equal(t, []int{}, in)
}

// 関数に引数として渡したSliceを関数内で操作しても
// 関数外には反映されない
// 値がポインタも同様
func TestSliceArgUpdatePointer(t *testing.T) {
	f := func(s []*int) {
		one := 1
		s = append(s, &one)
	}
	in := []*int{}
	f(in)
	assert.Equal(t, []*int{}, in)
}

// 関数に引数として渡したSliceを関数内で操作しても
// 関数外には反映されない
func TestMapArgUpdate(t *testing.T) {
	f := func(s map[int]int) {
		s[1] = 1
	}
	in := map[int]int{}
	f(in)
	assert.Equal(t, map[int]int{1: 1}, in)
}
