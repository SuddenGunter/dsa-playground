package solver_test

import (
	"testing"

	"github.com/SuddenGunter/dsa-playground/leetcode/1275/solver"
	"github.com/stretchr/testify/assert"
)

func TestSolve_Example1(t *testing.T) {
	input := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}

	result := solver.Solve(input)

	assert.Equal(t, "A", result)
}

func TestSolve_Example2(t *testing.T) {
	input := [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}

	result := solver.Solve(input)

	assert.Equal(t, "B", result)
}

func TestSolve_Example3(t *testing.T) {
	input := [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}

	result := solver.Solve(input)

	assert.Equal(t, "Draw", result)
}

func TestSolve_Example4(t *testing.T) {
	input := [][]int{{0, 0}, {1, 1}}

	result := solver.Solve(input)

	assert.Equal(t, "Pending", result)
}

func TestSolve_Example5(t *testing.T) {
	input := [][]int{{1, 1}}

	result := solver.Solve(input)

	assert.Equal(t, "Pending", result)
}

func TestSolve_Example6(t *testing.T) {
	input := [][]int{{0, 0}, {0, 1}, {2, 0}, {1, 0}, {2, 1}, {2, 2}, {1, 1}}

	result := solver.Solve(input)

	assert.Equal(t, "Pending", result)
}
