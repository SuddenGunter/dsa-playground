package main

type Stack[T any] struct {
	storage []T
	pos     int
}

func NewStack[T any](num int) *Stack[T] {
	return &Stack[T]{
		storage: make([]T, num),
		pos:     0,
	}
}

func (fs *Stack[T]) Push(val T) {
	if fs.pos == len(fs.storage) {
		fs.resize(2 * len(fs.storage))
	}

	fs.storage[fs.pos] = val
	fs.pos++
}

func (fs *Stack[T]) Pop() T {
	fs.pos--
	item := fs.storage[fs.pos]
	// if (N > 0 && N == a.length/4) resize(a.length/2);
	if fs.pos > 0 && fs.pos == len(fs.storage)/4 {
		fs.resize(len(fs.storage) / 2)
	}

	return item
}

func (fs *Stack[T]) Empty() bool {
	return fs.pos == 0
}

func (fs *Stack[T]) Size() int {
	return fs.pos
}

func (fs *Stack[T]) Iteartor() Iterator[T] {
	return &reverseSliceIterator[T]{
		slice: fs.storage,
		pos:   fs.pos - 1,
	}
}

// I know we do not need all this in go and could just use append/copy for slice.
// I use it cause I wanted to be as close as possible to book implementation
func (fs *Stack[T]) resize(max int) {
	old := fs.storage
	fs.storage = make([]T, max)
	for i := range old {
		fs.storage[i] = old[i]
	}
}

type reverseSliceIterator[T any] struct {
	slice []T
	pos   int
}

func (rsi *reverseSliceIterator[T]) HasNext() bool {
	return rsi.pos >= 0
}

func (rsi *reverseSliceIterator[T]) Next() T {
	item := rsi.slice[rsi.pos]
	rsi.pos--
	return item
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
