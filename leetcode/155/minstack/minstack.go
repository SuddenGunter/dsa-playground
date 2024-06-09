package minstack

import (
	"math"

	"github.com/SuddenGunter/dsa-playground/leetcode/155/linkedlist"
)

// MinStack implemented on top of the doubly linked list.
type MinStack struct {
	minimums *linkedlist.DoublyLinkedList
	head     *linkedlist.DoublyLinkedListNode
}

func Constructor() MinStack {
	minimums := &linkedlist.DoublyLinkedList{}
	minimums.Prepend(math.MaxInt)
	return MinStack{
		minimums: minimums,
	}
}

func (l *MinStack) Push(val int) {
	if val <= l.minimums.Top() {
		l.minimums.Prepend(val)
	}

	current := l.head
	if current == nil {
		l.head = &linkedlist.DoublyLinkedListNode{
			Key: val,
		}

		return
	}

	appended := &linkedlist.DoublyLinkedListNode{
		Key:  val,
		Next: current,
	}

	current.Prev = appended

	l.head = appended
}

func (l *MinStack) Pop() {
	if l.head.Key == l.minimums.Top() {
		l.minimums.Pop()
	}

	l.head = l.head.Next

	if l.head != nil {
		l.head.Prev = nil
	}
}

func (l *MinStack) Top() int {
	return l.head.Key
}

func (l *MinStack) GetMin() int {
	return l.minimums.Top()
}
