package solution_test

import (
	"testing"

	solution "github.com/SuddenGunter/dsa-playground/leetcode/622"
	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	myCircularQueue := solution.Constructor(3)
	assert.True(t, myCircularQueue.EnQueue(1))
	assert.True(t, myCircularQueue.EnQueue(2))
	assert.True(t, myCircularQueue.EnQueue(3))
	assert.False(t, myCircularQueue.EnQueue(4))
	assert.Equal(t, 3, myCircularQueue.Rear())
	assert.True(t, myCircularQueue.IsFull())
	assert.True(t, myCircularQueue.DeQueue())
	assert.True(t, myCircularQueue.EnQueue(4))
	assert.Equal(t, 4, myCircularQueue.Rear())
}
