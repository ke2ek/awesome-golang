package test

import (
	"awesome-golang/common"
	"awesome-golang/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	h := heap.NewIntHeap()
	data := []int{3, 4, 6, 1, 2, 5}
	for _, num := range data {
		heap.Push(h, num)
	}

	for i := 1; i < 7; i++ {
		val := heap.Pop(h).(int)
		assert.Equal(t, val, i, common.ERROR_MSG)
	}
}
