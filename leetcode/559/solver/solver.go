package solver

import "math"

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Children == nil || len(root.Children) == 0 {
		return 1
	}

	stack := stack{data: make([]stackEntry, 0, 1)}
	stack.push(stackEntry{
		Node:  root,
		Depth: 1,
	})

	max := math.MinInt16

	for !stack.isEmpty() {
		node := stack.pop()
		max = maxInt(max, node.Depth)

		for _, n := range node.Node.Children {
			stack.push(stackEntry{
				Node:  n,
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

func Solve(root *Node) int {
	return maxDepth(root)
}
