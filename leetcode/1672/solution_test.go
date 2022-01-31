package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumWealth_Example1(t *testing.T) {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}

	output := maximumWealth(accounts)

	assert.Equal(t, 6, output)
}

func TestMaximumWealth_Example2(t *testing.T) {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}

	output := maximumWealth(accounts)

	assert.Equal(t, 10, output)
}

func TestMaximumWealth_Example3(t *testing.T) {
	accounts := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}

	output := maximumWealth(accounts)

	assert.Equal(t, 17, output)
}
