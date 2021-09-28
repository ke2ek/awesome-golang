package test

import (
	"awesome-golang/common"
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
	assert.Equal(t, q.Size(), 5, common.ERROR_MSG)
	q.Pop()
	assert.Equal(t, q.Front(), "element2", common.ERROR_MSG)
	assert.Equal(t, q.Size(), 4, common.ERROR_MSG)
	q.Push("element6")
	assert.Equal(t, q.Back(), "element6", common.ERROR_MSG)
	q.Pop()
	q.Pop()
	assert.Equal(t, q.Front(), "element4", common.ERROR_MSG)
	assert.Equal(t, q.Size(), 3, common.ERROR_MSG)
	q.Pop()
	q.Pop()
	q.Pop()
	assert.Equal(t, q.Empty(), true, common.ERROR_MSG)
}
