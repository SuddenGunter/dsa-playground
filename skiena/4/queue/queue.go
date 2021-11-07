package queue

type Element struct {
	Data float64
	next *Element
}

type Queue struct {
	head *Element
	last *Element
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		last: nil,
	}
}

func (q *Queue) Enqueue(e Element) {
	if q.IsEmpty() {
		q.head = &e
		q.last = &e
		return
	}

	q.last.next = &e
	q.last = q.last.next
}

func (q *Queue) Dequeue() *Element {
	if q.IsEmpty() {
		return nil
	}

	tmp := q.head
	q.head = q.head.next

	if q.head == nil {
		q.last = nil
	}

	return tmp
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Peek() *Element {
	return q.head
}
