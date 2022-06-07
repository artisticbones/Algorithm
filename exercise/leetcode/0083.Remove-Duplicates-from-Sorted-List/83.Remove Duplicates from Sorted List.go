package _083_Remove_Duplicates_from_Sorted_List

import "github.com/artisticbones/Algorithm/exercise/leetcode/structures"

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if cur == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
