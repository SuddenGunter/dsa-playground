package solver

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Build(parents []int) *TreeNode {
	return addFromArray(parents, nil, 0)
}

func addFromArray(arr []int, t *TreeNode, i int) *TreeNode {
	if i < len(arr) && arr[i] != math.MinInt {

		tmp := TreeNode{
			Val:   arr[i],
			Left:  nil,
			Right: nil,
		}

		left := addFromArray(arr, tmp.Left, 2*i+1)
		right := addFromArray(arr, tmp.Right, 2*i+2)

		tmp.Left = left
		tmp.Right = right

		t = &tmp
	}

	return t
}
