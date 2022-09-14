/*
 * @lc app=leetcode.cn id=1619 lang=golang
 *
 * [1619] 删除某些元素后的数组均值
 */
package leetcode

import "sort"

// @lc code=start
func trimMean(arr []int) float64 {
	sum := 0
	sort.Ints(arr)
	for _, v := range arr[len(arr)/20 : len(arr)*19/20] {
		sum += v
	}
	return float64(sum*10) / float64(len(arr)*9)
}

// @lc code=end
