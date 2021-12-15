package solver_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SuddenGunter/dsa-playground/leetcode/559/solver"
)

func TestExample1(t *testing.T) {
	root := solver.Build([]int{1, math.MinInt, 3, 2, 4, math.MinInt, 5, 6})

	result := solver.Solve(root)

	assert.Equal(t, 3, result)
}

func TestExample2(t *testing.T) {
	root := solver.Build([]int{1, math.MinInt, 2, 3, 4, 5, math.MinInt, math.MinInt, 6, 7, math.MinInt, 8, math.MinInt, 9, 10, math.MinInt, math.MinInt, 11, math.MinInt, 12, math.MinInt, 13, math.MinInt, math.MinInt, 14})

	result := solver.Solve(root)

	assert.Equal(t, 5, result)
}

func TestExample3(t *testing.T) {
	root := solver.Build([]int{})

	result := solver.Solve(root)

	assert.Equal(t, 0, result)
}

func TestExample4(t *testing.T) {
	root := solver.Build([]int{0})

	result := solver.Solve(root)

	assert.Equal(t, 1, result)
}
