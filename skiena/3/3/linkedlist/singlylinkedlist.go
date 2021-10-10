package linkedlist

type SinglyLinkedListNode struct {
	key  int
	next *SinglyLinkedListNode
}

func (n *SinglyLinkedListNode) GetKey() int {
	return n.key
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

// Append to a linked list.
// O(n).
func (l *SinglyLinkedList) Append(key int) {
	current := l.head
	if current == nil {
		l.head = &SinglyLinkedListNode{
			key: key,
		}

		return
	}

	// O(n)
	for next := current.next; next != nil; next = current.next {
		current = current.next
	}

	current.next = &SinglyLinkedListNode{
		key: key,
	}
}

func (l *SinglyLinkedList) Foreach(f func(n *SinglyLinkedListNode)) {
	current := l.head
	if current == nil {
		return
	}

	for current != nil {
		f(current)
		current = current.next
	}
}

func (l *SinglyLinkedList) Reverse() {
	current := l.head
	if current == nil {
		return
	}

	var prev *SinglyLinkedListNode

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev
}
