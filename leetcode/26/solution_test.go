package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates_Example1(t *testing.T) {
	nums := []int{1, 1, 2}
	expectedNums := []int{1, 2}

	k := removeDuplicates(nums)

	assert.Equal(t, len(expectedNums), k)
	for i := 0; i < k; i++ {
		assert.Equal(t, expectedNums[i], nums[i])
	}
}

func TestTwoSum_Example2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums := []int{0, 1, 2, 3, 4}

	k := removeDuplicates(nums)

	assert.Equal(t, len(expectedNums), k)
	for i := 0; i < k; i++ {
		assert.Equal(t, expectedNums[i], nums[i])
	}
}

func TestTwoSum_Example3(t *testing.T) {
	nums := []int{0, 1, 2}
	expectedNums := []int{0, 1, 2}

	k := removeDuplicates(nums)

	assert.Equal(t, len(expectedNums), k)
	for i := 0; i < k; i++ {
		assert.Equal(t, expectedNums[i], nums[i])
	}
}
