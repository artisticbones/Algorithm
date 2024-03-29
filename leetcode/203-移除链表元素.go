package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 处理头节点
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	p, q := head, head.Next
	for q != nil {
		if q.Val == val {
			p.Next, q = q.Next, q.Next
		} else {
			p, q = p.Next, q.Next
		}
	}

	return head
}

func removeElements_2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements_2(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
