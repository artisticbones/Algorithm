package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headB == nil || headA == nil {
//		return nil
//	}
//
//	curA := headA
//	for curA != nil {
//		curB := headB
//		for curB != nil {
//			if curA == curB {
//				return curA
//			}
//			curB = curB.Next
//		}
//		curA = curA.Next
//	}
//	return nil
//}

// getIntersectionNode 双指针法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headB == nil || headA == nil {
		return nil
	}

	curA, curB := headA, headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}

	return curA
}
