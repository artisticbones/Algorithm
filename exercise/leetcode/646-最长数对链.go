/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */
package leetcode

import (
	"sort"
)

// @lc code=start
//
//	func findLongestChain(pairs [][]int) int {
//		sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
//		dp := make([]int, len(pairs))
//		max := func(a, b int) int {
//			if b > a {
//				return b
//			}
//			return a
//		}
//		for i, p := range pairs {
//			dp[i] = 1
//			for j, q := range pairs[:i] {
//				if p[0] > q[1] {
//					dp[i] = max(dp[i], dp[j]+1)
//				}
//			}
//		}
//		return dp[len(pairs)-1]
//	}
//
// 贪心算法
// 要挑选最长数对链的第一个数对时，最优的选择是挑选第二个数字最小的，这样能给挑选后续的数对留下更多的空间。
// 挑完第一个数对后，要挑第二个数对时，也是按照相同的思路，是在剩下的数对中，第一个数字满足题意的条件下，挑选第二个数字最小的。
// 按照这样的思路，可以先将输入按照第二个数字排序，然后不停地判断第一个数字是否能满足大于前一个数对的第二个数字即可。
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	cur := pairs[0][1]
	ans := 1
	for _, v := range pairs {
		if cur < v[0] {
			cur = v[1]
			ans++
		}
	}
	return ans
}

// @lc code=end
