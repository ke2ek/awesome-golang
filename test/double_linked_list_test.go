package test

import (
	"awesome-golang/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func appendToDll(l *list.DoubleLinkedList, data []int) {
	for _, num := range data {
		l.PushBack(num)
	}
}

func validateDll(t *testing.T, l *list.DoubleLinkedList, data []int) {
	node := l.First()
	for _, num := range data {
		assert.Equal(t, num, node.Value)
		node = node.Next()
	}
}

func TestDoubleLinkedList(t *testing.T) {
	dll := list.NewDoubleLinkedList()
	data := []int{7, 3, 6, 1, 0, 8, 1, 6}
	appendToDll(dll, data)
	dll.Iterate()
	dll.PushFront(0)
	dll.PushFront(1)
	dll.PushFront(0)
	dll.Iterate()
	assert.Equal(t, 11, dll.Size())
	for i := 0; i < 3; i++ {
		dll.PopBack()
		dll.PopFront()
	}
	assert.Equal(t, 5, dll.Size())
	ans := []int{7, 3, 6, 1, 0}
	validateDll(t, dll, ans)
	dll.Reverse()
	ans = []int{0, 1, 6, 3, 7}
	validateDll(t, dll, ans)
	dll.Iterate()
	data = []int{0, 0, 0, 0, 5, 8, 3, 0}
	appendToDll(dll, data)
	ans = []int{0, 1, 6, 3, 7, 0, 0, 0, 0, 5, 8, 3, 0}
	validateDll(t, dll, ans)
	dll.Unique()
	ans = []int{0, 1, 6, 3, 7, 5, 8}
	validateDll(t, dll, ans)

	dll2 := list.NewDoubleLinkedList()
	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	appendToDll(dll2, data)
	toNode := dll2.Find(5)
	dll2.Insert(toNode, dll)
	ans = []int{1, 2, 3, 4, 0, 1, 6, 3, 7, 5, 8, 5, 6, 7, 8, 9}
	validateDll(t, dll2, ans)
	assert.Equal(t, len(ans), dll2.Size())
	dll2.Unique()
	ans = []int{1, 2, 3, 4, 0, 6, 7, 5, 8, 9}
	validateDll(t, dll2, ans)
	assert.Equal(t, len(ans), dll2.Size())
	dll2.Erase(0)
	dll2.Erase(5)
	ans = []int{1, 2, 3, 4, 6, 7, 8, 9}
	validateDll(t, dll2, ans)
	assert.Equal(t, len(ans), dll2.Size())
}
