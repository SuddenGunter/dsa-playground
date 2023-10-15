package main

type node[T any] struct {
	item T
	next *node[T]
}

type Stack[T any] struct {
	first *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	if s.first == nil {
		s.first = &node[T]{
			item: item,
			next: nil,
		}

		return
	}

	oldFirst := s.first
	s.first = &node[T]{
		item: item,
		next: oldFirst,
	}
}

func (s *Stack[T]) Pop() T {
	result := s.first.item
	s.first = s.first.next

	return result
}

func (s *Stack[T]) Empty() bool {
	return s.first == nil
}

func (s *Stack[T]) Size() int {
	size := 0
	for current := s.first; current != nil; current = current.next {
		size++
	}

	return size
}

func (s Stack[T]) Peek() T {
	return s.first.item
}

func (s *Stack[T]) Iteartor() Iterator[T] {
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
