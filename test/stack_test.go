package test

import (
	"awesome-golang/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stk := stack.New()
	stk.Push(123)
	stk.Push(489)
	stk.Push(973)
	assert.Equal(t, 3, stk.Size())
	assert.Equal(t, 973, stk.Top())
	stk.Pop()
	assert.Equal(t, 489, stk.Top())
	stk.Pop()
	assert.Equal(t, 123, stk.Top())
	stk.Pop()
	assert.Equal(t, true, stk.Empty())
}
