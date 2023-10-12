package main

type node[T any] struct {
	item T
	next *node[T]
}

type Queue[T any] struct {
	first *node[T]
	last  *node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (s *Queue[T]) Enqueue(item T) {
	if s.first == nil {
		s.first = &node[T]{
			item: item,
			next: nil,
		}

		s.last = s.first

		return
	}

	s.last.next = &node[T]{
		item: item,
		next: nil,
	}

	s.last = s.last.next
}

func (s *Queue[T]) Dequeue() T {
	if s.first == nil {
		var zero T
		return zero
	}

	result := s.first.item
	s.first = s.first.next

	return result
}

func (s *Queue[T]) Empty() bool {
	return s.first == nil
}

func (s *Queue[T]) Size() int {
	size := 0
	for current := s.first; current != nil; current = current.next {
		size++
	}

	return size
}

func (s *Queue[T]) Iteartor() Iterator[T] {
	return &linkedListIterator[T]{
		next: s.first,
	}
}

type linkedListIterator[T any] struct {
	next *node[T]
}

func (lli *linkedListIterator[T]) HasNext() bool {
	return lli.next != nil
}

func (lli *linkedListIterator[T]) Next() T {
	item := lli.next.item
	lli.next = lli.next.next
	return item
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
