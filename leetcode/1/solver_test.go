package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum_Example1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	ouput := twoSum(nums, target)

	assert.Equal(t, []int{0, 1}, ouput)
}

func TestTwoSum_Example2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	ouput := twoSum(nums, target)

	assert.Equal(t, []int{1, 2}, ouput)
}

func TestTwoSum_Example3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	ouput := twoSum(nums, target)

	assert.Equal(t, []int{0, 1}, ouput)
}

func TestTwoSum_Example4(t *testing.T) {
	nums := []int{0, 0}
	target := 0

	ouput := twoSum(nums, target)

	assert.Equal(t, []int{0, 1}, ouput)
}

func TestTwoSum_Example5(t *testing.T) {
	nums := []int{5, -5}
	target := 0

	ouput := twoSum(nums, target)

	assert.Equal(t, []int{0, 1}, ouput)
}
