package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func fromSlice(x []int) *TreeNode {
	return insertLevelOrder(x, nil, 0)
}

// insertLevelOrder inserts nodes into the binary tree in level order
func insertLevelOrder(arr []int, root *TreeNode, i int) *TreeNode {
	// Base case for recursion
	if i < len(arr) {
		temp := &TreeNode{Val: arr[i]}
		root = temp

		// Insert left child
		root.Left = insertLevelOrder(arr, root.Left, 2*i+1)

		// Insert right child
		root.Right = insertLevelOrder(arr, root.Right, 2*i+2)
	}
	return root
}

func asSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.Val)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}
