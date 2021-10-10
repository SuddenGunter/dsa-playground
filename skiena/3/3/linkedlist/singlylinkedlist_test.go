package linkedlist_test

import (
	"testing"

	"github.com/SuddenGunter/dsa-playground/skiena/3/3/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestAppendWorksAsExpected(t *testing.T) {
	list := &linkedlist.SinglyLinkedList{}
	expectedValues := []int{1, 2, 3}

	for _, v := range expectedValues {
		list.Append(v)
	}

	realValues := make([]int, 0)
	list.Foreach(func(n *linkedlist.SinglyLinkedListNode) {
		realValues = append(realValues, n.GetKey())
	})

	assert.Equal(t, expectedValues, realValues)
}

func TestReverseWorksAsExpected(t *testing.T) {
	list := &linkedlist.SinglyLinkedList{}
	initValues := []int{1, 2, 3}
	expectedValues := []int{3, 2, 1}

	for _, v := range initValues {
		list.Append(v)
	}

	list.Reverse()

	realValues := make([]int, 0)
	list.Foreach(func(n *linkedlist.SinglyLinkedListNode) {
		realValues = append(realValues, n.GetKey())
	})

	assert.Equal(t, expectedValues, realValues)
}
