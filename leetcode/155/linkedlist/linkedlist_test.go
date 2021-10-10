package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SuddenGunter/dsa-playground/leetcode/155/linkedlist"
	"github.com/stretchr/testify/require"
)

// TestPrependWorksAsExpected tests that values are being inserted in reverse order and all public methods are working.
func TestPrependWorksAsExpected(t *testing.T) {
	list := &linkedlist.DoublyLinkedList{}
	insertedValues := []int{1, 2, 3}
	for _, v := range insertedValues {
		list.Prepend(v)
	}

	empty := list.Empty()
	require.False(t, empty)

	top := list.Top()
	require.Equal(t, 3, top)

	expectedValues := []int{3, 2, 1}
	realValues := make([]int, 0)
	for !list.Empty() {
		realValues = append(realValues, list.TakeTop())
	}

	assert.Equal(t, expectedValues, realValues)
}
