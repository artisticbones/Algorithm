/*
 * @lc app=leetcode.cn id=1608 lang=golang
 *
 * [1608] 特殊数组的特征值
 */
package leetcode

import "sort"

// @lc code=start
func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	n := len(nums)
	for i := 1; i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}

// @lc code=end
