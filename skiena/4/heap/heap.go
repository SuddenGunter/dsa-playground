package heap

import "errors"

var minDominatesComparator = func(a, b float64) int8 {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

type Heap struct {
	body       []float64
	comparator func(a, b float64) int8
}

func NewHeap() *Heap {
	return &Heap{
		body:       []float64{},
		comparator: minDominatesComparator,
	}
}

func FromSlice(data []float64) *Heap {
	h := &Heap{
		body:       make([]float64, 0, cap(data)),
		comparator: minDominatesComparator,
	}

	for _, v := range data {
		h.Insert(v)
	}

	return h
}

func FromSliceFast(data []float64) *Heap {
	h := &Heap{
		body:       make([]float64, cap(data), cap(data)),
		comparator: minDominatesComparator,
	}

	for i, v := range data {
		h.body[i] = v
	}

	for i := len(h.body)/2 - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}

	return h
}

func Heapsort(data []float64) []float64 {
	buffer := make([]float64, 0, cap(data))
	h := FromSlice(data)
	for i := 0; i < cap(buffer); i++ {
		top, _ := h.TakeTop()
		buffer = append(buffer, top)
	}
	return buffer
}

func (h *Heap) parentIndexOf(child int) int {
	child += 1

	if child == 1 {
		return -1
	}

	return (child / 2) - 1
}

func (h *Heap) leftChildIndexOf(parent int) int {
	return 2*(parent+1) - 1
}

func (h *Heap) Insert(x float64) {
	h.body = append(h.body, x)
	h.bubbleUp(len(h.body) - 1)
}

func (h *Heap) bubbleUp(index int) {
	parentIndex := h.parentIndexOf(index)

	if parentIndex == -1 {
		return
	}

	if h.comparator(h.body[parentIndex], h.body[index]) == 1 {
		h.body[index], h.body[parentIndex] = h.body[parentIndex], h.body[index]
		h.bubbleUp(parentIndex)
	}
}

func (h *Heap) TakeTop() (float64, error) {
	if len(h.body) == 0 {
		return 0, errors.New("empty heap")
	}

	top := h.body[0]
	h.body[0] = h.body[len(h.body)-1]
	h.body = h.body[:len(h.body)-1]
	h.bubbleDown(0)

	return top, nil
}

func (h *Heap) bubbleDown(index int) {
	leftChildIndex := h.leftChildIndexOf(index)

	topIndex := index

	// check that h.body[index] dominates both left and right children
	for i := 0; i <= 1; i++ {
		if leftChildIndex+i <= len(h.body)-1 {
			if h.comparator(h.body[topIndex], h.body[leftChildIndex+i]) == 1 {
				topIndex = leftChildIndex + i
			}
		}
	}

	// if h.body[index] doesn't dominate on child - swap and check the same for the lower level
	if topIndex != index {
		h.body[index], h.body[topIndex] = h.body[topIndex], h.body[index]
		h.bubbleDown(topIndex)
	}
}
