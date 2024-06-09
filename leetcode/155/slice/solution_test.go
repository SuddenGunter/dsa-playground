package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin()
	minStack.Pop()
	top := minStack.Top()
	min := minStack.GetMin()

	assert.Equal(t, 0, top)
	assert.Equal(t, -2, min)
}
