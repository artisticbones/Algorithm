package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	stack, p := make([]int, 0, 1), head
//
//	for p != nil {
//		stack = append(stack, p.Val)
//		p = p.Next
//	}
//
//	var ret = &ListNode{Val: stack[len(stack)-1], Next: nil}
//	p = ret
//	for i := len(stack) - 2; i >= 0; i-- {
//		q := &ListNode{Val: stack[i], Next: nil}
//		p.Next = q
//		p = p.Next
//	}
//	return ret
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
