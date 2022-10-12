/*
 * @lc app=leetcode.cn id=817 lang=golang
 *
 * [817] 链表组件
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func numComponents(head *ListNode, nums []int) int {
	if head == nil || len(nums) == 0 {
		return 0
	}
	ans := 0
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}
	for inSet := false; head != nil; head = head.Next {
		if _, ok := set[head.Val]; !ok {
			inSet = false
		} else if !inSet {
			inSet = true
			ans++
		}
	}
	return ans
}

// @lc code=end
