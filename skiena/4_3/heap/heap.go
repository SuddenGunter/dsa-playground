package heap

type Heap struct {
	body       []rune
	comparator func(a, b rune) int8
}

func NewHeap() *Heap {
	return &Heap{
		body: []rune{},
		comparator: func(a, b rune) int8 {
			switch {
			case a > b:
				return 1
			case a < b:
				return -1
			default:
				return 0
			}
		},
	}
}

func FromSlice(data []rune) *Heap {
	h := NewHeap()
	for _, v := range data {
		h.Insert(v)
	}

	return h
}

// todo: issue because indexes starts from 1
func (h *Heap) parentIndexOf(child int) int {
	if child == 1 {
		return -1
	}

	return child / 2
}

func (h *Heap) leftChildIndexOf(parent int) int {
	return 2 * parent
}

func (h *Heap) Insert(x rune) {
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
