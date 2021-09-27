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
	assert.Equal(t, stk.Size(), 3, "The two values should be the same.")
	assert.Equal(t, stk.Top(), 973, "The two values should be the same.")
	stk.Pop()
	assert.Equal(t, stk.Top(), 489, "The two values should be the same.")
	stk.Pop()
	assert.Equal(t, stk.Top(), 123, "The two values should be the same.")
	stk.Pop()
	assert.Equal(t, stk.Empty(), true, "The two values should be the same.")
}
