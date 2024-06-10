package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(t, 7, result)
}

func TestExample2(t *testing.T) {
	result := maxProfit([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 4, result)
}

func TestExample3(t *testing.T) {
	result := maxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(t, 0, result)
}
