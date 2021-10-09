package test

import (
	"awesome-golang/queue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := queue.New()
	q.Push("element1")
	q.Push("element2")
	q.Push("element3")
	q.Push("element4")
	q.Push("element5")
	assert.Equal(t, 5, q.Size())
	q.Pop()
	assert.Equal(t, "element2", q.Front())
	assert.Equal(t, 4, q.Size())
	q.Push("element6")
	assert.Equal(t, "element6", q.Back())
	q.Pop()
	q.Pop()
	assert.Equal(t, "element4", q.Front())
	assert.Equal(t, 3, q.Size())
	q.Pop()
	q.Pop()
	q.Pop()
	assert.Equal(t, true, q.Empty())
}
