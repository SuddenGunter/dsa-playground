package solver

//
//func maxDepth(root *Node) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//
//	stack := stack{data: make([]stackEntry, 0, 1)}
//	stack.push(stackEntry{
//		Node:  root,
//		Depth: 1,
//	})
//
//	max := math.MinInt16
//
//	for !stack.isEmpty() {
//		node := stack.pop()
//		max = maxInt(max, node.Depth)
//
//		if node.Node.Left != nil {
//			stack.push(stackEntry{
//				Node:  node.Node.Left,
//				Depth: node.Depth + 1,
//			})
//		}
//
//		if node.Node.Right != nil {
//			stack.push(stackEntry{
//				Node:  node.Node.Right,
//				Depth: node.Depth + 1,
//			})
//		}
//
//	}
//
//	return max
//}
//
//func maxInt(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}

func Solve(root *Node) int {
	return -1
}
