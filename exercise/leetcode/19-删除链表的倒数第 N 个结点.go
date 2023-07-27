package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Val: 0, Next: nil}

	ints := make([]int, 0, 1)
	p := head
	for p != nil {
		ints = append(ints, p.Val)
		p = p.Next
	}
	removeIndex := len(ints) - n
	cur := dummyHead
	for i, val := range ints {
		if i == removeIndex {
			continue
		}
		q := &ListNode{Val: val, Next: nil}
		cur.Next = q
		cur = cur.Next
	}

	return dummyHead.Next
}
