package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	result := coinChange([]int{1, 2, 5}, 11)
	assert.Equal(t, 3, result)
}

func TestExample2(t *testing.T) {
	result := coinChange([]int{2}, 3)
	assert.Equal(t, -1, result)
}

func TestExample3(t *testing.T) {
	result := coinChange([]int{1}, 0)
	assert.Equal(t, 0, result)
}
