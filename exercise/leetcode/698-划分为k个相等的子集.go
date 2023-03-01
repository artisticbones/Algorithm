/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */
package leetcode

import "sort"

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	all := 0
	for _, v := range nums {
		all += v
	}
	if all%k != 0 {
		return false
	}
	// 每个集合的平均值
	per := all / k
	sort.Ints(nums)
	if nums[len(nums)-1] > per {
		return false
	}
	dp := make([]bool, 1<<len(nums))
	dp[0] = true
	curSum := make([]int, 1<<len(nums))
	for i, v := range dp {
		if !v {
			continue
		}
		for j, num := range nums {
			if curSum[i]+num > per {
				break
			}
			if i>>j&1 == 0 {
				next := i | 1<<j
				if !dp[next] {
					curSum[next] = (curSum[i] + nums[j]) % per
					dp[next] = true
				}
			}
		}
	}
	return dp[1<<len(nums)-1]
}

// @lc code=end
