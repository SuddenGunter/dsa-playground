package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(t, 5, result)
}

func TestExample2(t *testing.T) {
	result := maxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(t, 0, result)
}
