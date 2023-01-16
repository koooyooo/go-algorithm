package _basic_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type val struct {
	v int
}

var createValMap = func() map[string]val {
	m := make(map[string]val)
	m["key"] = val{2}
	return m
}

var createPointerMap = func() map[string]*val {
	m := make(map[string]*val)
	for k, v := range createValMap() {
		m[k] = &v
	}
	return m
}

func TestEmptyVal(t *testing.T) {
	m := make(map[string]val)
	v, ok := m["nonexistent"]
	// ok フラグは偽を返す
	assert.False(t, ok)
	// 値はゼロ値が返るが、mapの中身を返している訳ではない
	assert.Equal(t, val{}, v)
}

func TestGetMapStruct(t *testing.T) {
	// 関数からの戻りの mapは値も保持している（コピーだけどね）
	m := createValMap()
	assert.Equal(t, 2, m["key"].v)
}

func TestUpdateMapVal(t *testing.T) {
	// valueが値の形式だと mapの外で更新できない（コピー値だからね）
	m := createValMap()
	val := m["key"]
	val.v = 3
	assert.Equal(t, 2, m["key"].v)
}

func TestUpdateMapPointer(t *testing.T) {
	// valueが値の形式だと mapの外で更新できる（参照だからね）
	m := createPointerMap()
	val := m["key"]
	val.v = 3
	assert.Equal(t, 3, m["key"].v)
}

// ゼロ値があれば初期化が不要と思い mapのvalueに値を入れようと考えていたが、後からの操作が必要ならポインタ一択である
