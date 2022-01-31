package solution

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	current := &ListNode{
		Val:  -1,
		Next: nil,
	}

	head := &current.Next

	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			current.Next = &ListNode{
				Val:  p1.Val,
				Next: nil,
			}
			p1 = p1.Next
		} else {
			current.Next = &ListNode{
				Val:  p2.Val,
				Next: nil,
			}
			p2 = p2.Next
		}

		current = current.Next
	}

	nn := onlyOneNotNil(p1, p2)
	for nn != nil {
		next := &ListNode{
			Val:  nn.Val,
			Next: nil,
		}

		current.Next = next

		current = next
		nn = nn.Next
	}

	return *head
}

func onlyOneNotNil(p1 *ListNode, p2 *ListNode) *ListNode {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	return nil
}
