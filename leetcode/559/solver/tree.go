package solver

import "math"

type Node struct {
	Val      int
	Children []*Node
}

func Build(nodes []int) *Node {
	iterator := 0

	root := &Node{
		Val:      nodes[iterator],
		Children: nil,
	}

	s := queue{elements: make([]queueEntry, 0)}
	s.enqueue(queueEntry{
		Node: root,
	})

	iterator = 2

	for !s.isEmpty() && iterator < len(nodes) {
		children := make([]*Node, 0)
		val := nodes[iterator]

		node := s.dequeue().Node

		for val != math.MinInt && iterator < len(nodes) {
			child := &Node{
				Val:      val,
				Children: nil,
			}

			children = append(children, child)
			s.enqueue(queueEntry{
				Node: child,
			})

			iterator++

			if iterator >= len(nodes) {
				break
			}

			val = nodes[iterator]
		}

		iterator++

		node.Children = children
	}

	return root
}

type queueEntry struct {
	Node *Node
}

type queue struct {
	elements []queueEntry
}

func (q *queue) enqueue(p queueEntry) {
	q.elements = append(q.elements, p)
}

func (q *queue) dequeue() queueEntry {
	returned := q.elements[0]

	q.elements = q.elements[1:]

	return returned
}

func (q *queue) isEmpty() bool {
	return len(q.elements) == 0
}
