/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */
package leetcode

import "sort"

// @lc code=start
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
	dp := make([]int, len(pairs))
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	for i, p := range pairs {
		dp[i] = 1
		for j, q := range pairs[:i] {
			if p[0] > q[1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(pairs)-1]
}

// @lc code=end
