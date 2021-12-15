package solver

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	queue := queue{elements: make([]queueEntry, 0, 1)}

	queue.enqueue(queueEntry{
		Node:  root,
		Depth: 1,
	})

	max := math.MinInt16

	for !queue.isEmpty() {
		node := queue.dequeue()
		max = maxInt(max, node.Depth)

		if node.Node.Left == nil && node.Node.Right == nil {
			break
		}

		if node.Node.Left != nil {
			queue.enqueue(queueEntry{
				Node:  node.Node.Left,
				Depth: node.Depth + 1,
			})
		}

		if node.Node.Right != nil {
			queue.enqueue(queueEntry{
				Node:  node.Node.Right,
				Depth: node.Depth + 1,
			})
		}
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Solve(root *TreeNode) int {
	return minDepth(root)
}
