package solver

type stackEntry struct {
	Node  *Node
	Depth int
}

type stack struct {
	data []stackEntry
}

func (s *stack) pop() stackEntry {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack) push(p stackEntry) {
	s.data = append(s.data, p)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() stackEntry {
	return s.data[len(s.data)-1]
}
