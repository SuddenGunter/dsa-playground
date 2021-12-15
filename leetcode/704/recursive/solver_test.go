package solver_test

import (
	"testing"

	"github.com/SuddenGunter/dsa-playground/leetcode/704/recursive"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	result := solver.Solve(nums, target)

	assert.Equal(t, 4, result)
}

func TestExample2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2

	result := solver.Solve(nums, target)

	assert.Equal(t, -1, result)
}

func TestExample3(t *testing.T) {
	nums := []int{}
	target := 2

	result := solver.Solve(nums, target)

	assert.Equal(t, -1, result)
}

func TestExample4(t *testing.T) {
	nums := []int{-1, 0, 1, 3, 5, 9, 12}
	target := 5

	result := solver.Solve(nums, target)

	assert.Equal(t, 4, result)
}

func TestExample5(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 13

	result := solver.Solve(nums, target)

	assert.Equal(t, -1, result)
}

func TestExample6(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 12

	result := solver.Solve(nums, target)

	assert.Equal(t, 5, result)
}

func TestExample7(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := -1

	result := solver.Solve(nums, target)

	assert.Equal(t, 0, result)
}
