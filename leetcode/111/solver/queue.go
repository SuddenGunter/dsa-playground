package solver

type queueEntry struct {
	Node  *TreeNode
	Depth int
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
