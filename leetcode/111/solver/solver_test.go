package solver_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SuddenGunter/dsa-playground/leetcode/111/solver"
)

func TestExample1(t *testing.T) {
	root := solver.Build([]int{3, 9, 20, math.MinInt, math.MinInt, 15, 7})

	result := solver.Solve(root)

	assert.Equal(t, 2, result)
}

func TestExample2(t *testing.T) {
	root := solver.Build([]int{})

	result := solver.Solve(root)

	assert.Equal(t, 0, result)
}

func TestExample3(t *testing.T) {
	root := solver.Build([]int{0})

	result := solver.Solve(root)

	assert.Equal(t, 1, result)
}
