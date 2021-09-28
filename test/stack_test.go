package test

import (
	"awesome-golang/common"
	"awesome-golang/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stk := stack.New()
	stk.Push(123)
	stk.Push(489)
	stk.Push(973)
	assert.Equal(t, stk.Size(), 3, common.ERROR_MSG)
	assert.Equal(t, stk.Top(), 973, common.ERROR_MSG)
	stk.Pop()
	assert.Equal(t, stk.Top(), 489, common.ERROR_MSG)
	stk.Pop()
	assert.Equal(t, stk.Top(), 123, common.ERROR_MSG)
	stk.Pop()
	assert.Equal(t, stk.Empty(), true, common.ERROR_MSG)
}
