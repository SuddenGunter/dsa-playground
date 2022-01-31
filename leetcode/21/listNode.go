package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromSlice(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	current := &ListNode{
		Val:  slice[0],
		Next: nil,
	}

	head := current

	for _, v := range slice[1:] {
		next := &ListNode{
			Val:  v,
			Next: nil,
		}

		current.Next = next

		current = next
	}

	return head
}
