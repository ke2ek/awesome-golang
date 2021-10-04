package test

import (
	"awesome-golang/common"
	"awesome-golang/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func appendToSll(l *list.SingleLinkedList, data []int) {
	for _, num := range data {
		l.PushBack(num)
	}
}

func validateSll(t *testing.T, l *list.SingleLinkedList, data []int) {
	node := l.First()
	for _, num := range data {
		assert.Equal(t, node.Value, num, common.ERROR_MSG)
		node = node.Next()
	}
}

func TestSingleLinkedList(t *testing.T) {
	sll := list.NewSingleLinkedList()
	data := []int{7, 3, 6, 1, 0, 8, 1, 6}
	appendToSll(sll, data)
	sll.Iterate()
	sll.PushFront(0)
	sll.PushFront(1)
	sll.PushFront(0)
	sll.Iterate()
	assert.Equal(t, sll.Size(), 11, common.ERROR_MSG)
	for i := 0; i < 3; i++ {
		sll.PopBack()
		sll.PopFront()
	}
	ans := []int{7, 3, 6, 1, 0}
	validateSll(t, sll, ans)
	sll.Reverse()
	ans = []int{0, 1, 6, 3, 7}
	validateSll(t, sll, ans)
	sll.Iterate()
	data = []int{0, 0, 0, 0, 5, 8, 3, 0}
	appendToSll(sll, data)
	ans = []int{0, 1, 6, 3, 7, 0, 0, 0, 0, 5, 8, 3, 0}
	validateSll(t, sll, ans)
	assert.Equal(t, sll.Size(), len(ans), common.ERROR_MSG)
	sll.Unique()
	ans = []int{0, 1, 6, 3, 7, 5, 8}
	validateSll(t, sll, ans)
	assert.Equal(t, sll.Size(), len(ans), common.ERROR_MSG)
}
