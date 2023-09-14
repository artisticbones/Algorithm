/*
 * @lc app=leetcode.cn id=1624 lang=golang
 *
 * [1624] 两个相同字符之间的最长子字符串
 */
package leetcode

// @lc code=start

func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	firstIndex := make([]int, 26)
	for i := range firstIndex {
		firstIndex[i] = -1
	}
	for i, c := range s {
		c -= 'a'
		if firstIndex[c] < 0 {
			firstIndex[c] = i
		} else {
			ans = max(ans, i-firstIndex[c]-1)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
