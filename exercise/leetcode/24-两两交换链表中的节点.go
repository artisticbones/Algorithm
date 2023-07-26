package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func swapPairs(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	dummyHead := &ListNode{Val: 0, Next: head}
//	cur := dummyHead
//	for cur.Next != nil && cur.Next.Next != nil {
//		tmp1 := cur.Next
//		tmp2 := cur.Next.Next.Next
//
//		cur.Next = cur.Next.Next
//		cur.Next.Next = tmp1
//		tmp1.Next = tmp2
//
//		cur = cur.Next.Next
//	}
//	return dummyHead.Next
//}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head

	return next
}
