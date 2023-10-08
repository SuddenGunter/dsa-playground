package main

type FixedCapacityStack[T any] struct {
	storage []T
	pos     int
}

func NewFixedCapacityStack[T any](num int) *FixedCapacityStack[T] {
	return &FixedCapacityStack[T]{
		storage: make([]T, num),
		pos:     0,
	}
}

func (fs *FixedCapacityStack[T]) Push(val T) {
	fs.storage[fs.pos] = val
	fs.pos++
}

func (fs *FixedCapacityStack[T]) Pop() T {
	fs.pos--
	return fs.storage[fs.pos]
}

func (fs *FixedCapacityStack[T]) Empty() bool {
	return fs.pos == 0
}

func (fs *FixedCapacityStack[T]) Size() int {
	return fs.pos
}
