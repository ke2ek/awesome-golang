package test

import (
	"awesome-golang/common"
	linkedlist "awesome-golang/linked-list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedList(t *testing.T) {
	dll := linkedlist.NewDoubleLinkedList()
	data := []int{7, 3, 6, 1, 0, 8, 1, 6}
	for _, num := range data {
		dll.PushBack(num)
	}
	dll.Iterate()
	dll.PushFront(0)
	dll.PushFront(1)
	dll.PushFront(0)
	dll.Iterate()
	assert.Equal(t, dll.Size(), 11, common.ERROR_MSG)
	for i := 0; i < 3; i++ {
		dll.PopBack()
		dll.PopFront()
		dll.Iterate()
	}
	assert.Equal(t, dll.Size(), 5, common.ERROR_MSG)
	// 7 3 6 1 0 -> 0 1 6 3 7
	dll.Reverse()
	dll.Iterate()
	for i := 0; i < 4; i++ {
		dll.PushBack(0)
	}
	dll.PushBack(5)
	dll.PushBack(8)
	dll.PushBack(3)
	dll.PushBack(0)
	dll.Iterate() // 0 1 6 3 7 0 0 0 0 5 8 3 0
	dll.Unique()
	fmt.Print("dll1: ")
	dll.Iterate() // 0 1 6 3 7 5 8
	assert.Equal(t, dll.Size(), 7, common.ERROR_MSG)

	dll2 := linkedlist.NewDoubleLinkedList()
	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, num := range data {
		dll2.PushBack(num)
	}
	fmt.Print("dll2: ")
	dll2.Iterate()
	toNode := dll2.Find(5)
	dll2.Insert(toNode, dll)
	fmt.Print("dll2: ")
	dll2.Iterate() // 1 2 3 4 0 1 6 3 7 5 8 5 6 7 8 9
	assert.Equal(t, dll2.Size(), 16, common.ERROR_MSG)
	dll2.Unique()
	dll2.Iterate() // 1 2 3 4 0 6 7 5 8 9
	assert.Equal(t, dll2.Size(), 10, common.ERROR_MSG)
	dll2.Erase(0)
	dll2.Erase(5)
	dll2.Iterate() // 1 2 3 4 6 7 8 9
	assert.Equal(t, dll2.Size(), 8, common.ERROR_MSG)

	data = []int{0, 5, 10}
	dll3 := linkedlist.NewDoubleLinkedList()
	for _, num := range data {
		dll3.PushBack(num)
	}
}
