package test

import (
	"awesome-golang/common"
	linkedlist "awesome-golang/linked-list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLinkedList(t *testing.T) {
	sll := linkedlist.NewSingleLinkedList()
	data := []int{7, 3, 6, 1, 0, 8, 1, 6}
	for _, num := range data {
		sll.PushBack(num)
	}
	sll.Iterate()
	sll.PushFront(0)
	sll.PushFront(1)
	sll.PushFront(0)
	sll.Iterate()
	assert.Equal(t, sll.Size(), 11, common.ERROR_MSG)
	for i := 0; i < 3; i++ {
		sll.PopBack()
		sll.PopFront()
		sll.Iterate()
	}
	assert.Equal(t, sll.Size(), 5, common.ERROR_MSG)
	// 7 3 6 1 0 -> 0 1 6 3 7
	sll.Reverse()
	sll.Iterate()
	for i := 0; i < 4; i++ {
		sll.PushBack(0)
	}
	sll.PushBack(5)
	sll.PushBack(8)
	sll.PushBack(3)
	sll.PushBack(0)
	sll.Iterate() // 0 1 6 3 7 0 0 0 0 5 8 3 0
	sll.Unique()
	sll.Iterate() // 0 1 6 3 7 5 8
	assert.Equal(t, sll.Size(), 7, common.ERROR_MSG)
}
