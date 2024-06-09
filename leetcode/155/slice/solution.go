package slice

import "math"

// MinStack implemented on top of the slice.
type MinStack struct {
	elements *stack
	mins     *stack
}

func Constructor() MinStack {
	mins := &stack{}
	mins.push(math.MaxInt)
	return MinStack{
		elements: &stack{},
		mins:     mins,
	}
}

func (ms *MinStack) Push(val int) {
	ms.elements.push(val)
	if val <= ms.mins.top() {
		ms.mins.push(val)
	}
}

func (ms *MinStack) Pop() {
	x := ms.elements.top()
	ms.elements.pop()
	if x == ms.mins.top() {
		ms.mins.pop()
	}
}

func (ms *MinStack) Top() int {
	return ms.elements.top()
}

func (ms *MinStack) GetMin() int {
	return ms.mins.top()
}

type stack struct {
	elements []int
}

func (s *stack) push(x int) {
	s.elements = append(s.elements, x)
}

func (s *stack) pop() {
	s.elements = s.elements[:len(s.elements)-1]
}

func (s *stack) top() int {
	return s.elements[len(s.elements)-1]
}
