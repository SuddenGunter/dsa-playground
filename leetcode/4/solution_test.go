package solution_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/SuddenGunter/dsa-playground/leetcode/4"
)

func TestFindMedianSortedArrays_Example1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	result := solution.FindMedianSortedArrays(nums1, nums2)

	assert.Equal(t, 2., result)
}

func TestFindMedianSortedArrays_Example2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	result := solution.FindMedianSortedArrays(nums1, nums2)

	assert.Equal(t, 2.5, result)
}

func TestFindMedianSortedArrays_Example3(t *testing.T) {
	nums1 := []int{0, 0}
	nums2 := []int{0, 0}

	result := solution.FindMedianSortedArrays(nums1, nums2)

	assert.Equal(t, 0., result)
}

func TestFindMedianSortedArrays_Example4(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1}

	result := solution.FindMedianSortedArrays(nums1, nums2)

	assert.Equal(t, 1., result)
}

func TestFindMedianSortedArrays_Example5(t *testing.T) {
	nums1 := []int{2}
	nums2 := []int{}

	result := solution.FindMedianSortedArrays(nums1, nums2)

	assert.Equal(t, 2., result)
}

func TestFindMedianSortedArrays_Example6(t *testing.T) {
	nums1 := []int{1, 3, 8, 9, 15}
	nums2 := []int{7, 11, 18, 19, 21, 25}

	result := solution.FindMedianSortedArrays(nums1, nums2)

	assert.Equal(t, 11., result)
}

func TestFindMedianSortedArrays_Example7(t *testing.T) {
	nums1 := []int{23, 26, 31, 35}
	nums2 := []int{3, 5, 7, 9, 11, 16}

	result := solution.FindMedianSortedArrays(nums1, nums2)

	assert.Equal(t, 13.5, result)
}
