package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 使用切片语法构造的切片，在操作时，会影响到原来的切片，因为在同一块内存
func TestCopySlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	sNew := s[:len(s)-1]
	sNew = append(sNew, 6)
	assert.Equal(t, []int{1, 2, 3, 4, 6}, s)
	assert.Equal(t, []int{1, 2, 3, 4, 6}, sNew)
}
