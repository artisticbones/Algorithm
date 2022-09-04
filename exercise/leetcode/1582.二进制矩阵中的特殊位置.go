/*
 * @lc app=leetcode.cn id=1582 lang=golang
 *
 * [1582] 二进制矩阵中的特殊位置
 */
package leetcode

// @lc code=start
func numSpecial(mat [][]int) int {
	rows := make([]int, len(mat))
	cols := make([]int, len(mat[0]))
	for i, row := range mat {
		for j, col := range row {
			rows[i] += col
			cols[j] += col
		}
	}
	ans := 0
	for i, row := range mat {
		for j, x := range row {
			if x == 1 && rows[i] == 1 && cols[j] == 1 {
				ans++
			}
		}
	}
	return ans
}

// @lc code=end
