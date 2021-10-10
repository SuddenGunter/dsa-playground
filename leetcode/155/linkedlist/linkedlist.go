package linkedlist

type DoublyLinkedListNode struct {
	Key  int
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}

func (n *DoublyLinkedListNode) GetKey() int {
	return n.Key
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
}

func (l *DoublyLinkedList) Prepend(key int) {
	current := l.head
	if current == nil {
		l.head = &DoublyLinkedListNode{
			Key: key,
		}

		return
	}

	appended := &DoublyLinkedListNode{
		Key:  key,
		Next: current,
	}

	current.Prev = appended

	l.head = appended
}

func (l *DoublyLinkedList) Empty() bool {
	return l.head == nil
}

func (l *DoublyLinkedList) Top() int {
	return l.head.Key
}

func (l *DoublyLinkedList) Pop() {
	l.head = l.head.Next

	if l.head != nil {
		l.head.Prev = nil
	}
}

func (l *DoublyLinkedList) TakeTop() int {
	val := l.head.Key

	l.head = l.head.Next

	if l.head != nil {
		l.head.Prev = nil
	}

	return val
}
