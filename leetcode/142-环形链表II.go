package leetcode

//func detectCycle(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	fast, slow := head, head
//	for {
//		if fast == nil || fast.Next == nil {
//			return nil
//		}
//		fast = fast.Next.Next
//		slow = slow.Next
//
//		if fast == slow {
//			break
//		}
//	}
//	fast = head
//	for fast != slow {
//		fast = fast.Next
//		slow = slow.Next
//	}
//
//	return slow
//}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	keys := make(map[*ListNode]struct{}, 8)
	for head != nil {
		if _, ok := keys[head]; !ok {
			keys[head] = struct{}{}
		} else {
			return head
		}
		head = head.Next
	}
	return nil
}
