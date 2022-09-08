/*
 * @lc app=leetcode.cn id=667 lang=golang
 *
 * [667] 优美的排列 II
 */
package leetcode

// @lc code=start
func constructArray(n int, k int) []int {
	answer := make([]int, 0, n)
	// 初始化
	for i := 1; i < n-k; i++ {
		answer = append(answer, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		answer = append(answer, i)
		if i != j {
			answer = append(answer, j)
		}
		j--
	}
	return answer
}

// @lc code=end
