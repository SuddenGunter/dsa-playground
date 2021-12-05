package solution_test

import (
	"testing"

	solution "github.com/SuddenGunter/dsa-playground/leetcode/53/kadane"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArraySum_Example1(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	result := solution.MaxSubArraySum(nums)

	assert.Equal(t, 6, result)
}

func TestMaxSubArraySum_Example2(t *testing.T) {
	nums := []int{1}

	result := solution.MaxSubArraySum(nums)

	assert.Equal(t, 1, result)
}

func TestMaxSubArraySum_Example3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}

	result := solution.MaxSubArraySum(nums)

	assert.Equal(t, 23, result)
}

func TestMaxSubArraySum_Example4(t *testing.T) {
	nums := []int{-1}

	result := solution.MaxSubArraySum(nums)

	assert.Equal(t, -1, result)
}
