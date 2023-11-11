package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pHead := new(ListNode)
	current := pHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 == nil {
		current.Next = l2
		return pHead.Next
	}
	current.Next = l1
	return pHead.Next
}

//func test() {
//	l13 := ListNode{
//		Val:  9,
//		Next: nil,
//	}
//	l12 := ListNode{
//		Val:  4,
//		Next: &l13,
//	}
//	l1 := ListNode{
//		Val:  2,
//		Next: &l12,
//	}
//	l24 := ListNode{
//		Val:  9,
//		Next: nil,
//	}
//	l23 := ListNode{
//		Val:  6,
//		Next: &l24,
//	}
//	l22 := ListNode{
//		Val:  4,
//		Next: &l23,
//	}
//	l2 := ListNode{
//		Val:  1,
//		Next: &l22,
//	}
//	fmt.Println(mergeTwoLists(&l1, &l2))
//}
