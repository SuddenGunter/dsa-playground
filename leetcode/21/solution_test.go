package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists_Example1(t *testing.T) {
	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}
	expectedOutput := FromSlice([]int{1, 1, 2, 3, 4, 4})

	result := mergeTwoLists(FromSlice(list1), FromSlice(list2))

	assert.Equal(t, expectedOutput, result)
}

func TestMergeTwoLists_Example2(t *testing.T) {
	list1 := []int{}
	list2 := []int{}
	expectedOutput := FromSlice([]int{})

	result := mergeTwoLists(FromSlice(list1), FromSlice(list2))

	assert.Equal(t, expectedOutput, result)
}

func TestMergeTwoLists_Example3(t *testing.T) {
	list1 := []int{}
	list2 := []int{0}
	expectedOutput := FromSlice([]int{0})

	result := mergeTwoLists(FromSlice(list1), FromSlice(list2))

	assert.Equal(t, expectedOutput, result)
}

func TestMergeTwoLists_Example4(t *testing.T) {
	list1 := []int{0}
	list2 := []int{}
	expectedOutput := FromSlice([]int{0})

	result := mergeTwoLists(FromSlice(list1), FromSlice(list2))

	assert.Equal(t, expectedOutput, result)
}
