/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */
package leetcode

import "strconv"

// @lc code=start
//
//	func maximumSwap(num int) int {
//		n := make([]int, 0, 1)
//		var sum int
//		var swap func(i, j int)
//		var findPos func(array []int) (bigPos int, pos1 int, pos2 int)
//		swap = func(i, j int) { n[i], n[j] = n[j], n[i] }
//		findPos = func(array []int) (bigPos int, pos1 int, pos2 int) {
//			bigPos, pos1, pos2 = 0, len(array), len(array)
//			for i := 0; i < len(array); i++ {
//				if array[i] > array[bigPos] {
//					bigPos = i
//				} else if array[i] < array[bigPos] {
//					pos1, pos2 = i, bigPos
//				}
//			}
//			return
//		}
//		for num > 0 {
//			n = append(n, num%10)
//			num /= 10
//		}
//		_, pos1, pos2 := findPos(n)
//		if pos1 > len(n)-1 {
//			return num
//		}
//		swap(pos1, pos2)
//		for i := len(n) - 1; i >= 0; i-- {
//			sum = sum*10 + n[i]
//		}
//		return sum
//	}
func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)
	maxIdx, idx1, idx2 := n-1, -1, -1
	for i := n - 1; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			idx1, idx2 = i, maxIdx
		}
	}
	if idx1 < 0 {
		return num
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]
	v, _ := strconv.Atoi(string(s))
	return v
}

// @lc code=end
