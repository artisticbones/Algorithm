/*
 * @lc app=leetcode.cn id=828 lang=golang
 *
 * [828] 统计子串中的唯一字符
 */
package leetcode

// @lc code=start
func uniqueLetterString(s string) (ans int) {
	idx := map[rune][]int{}
	for i, c := range s {
		idx[c] = append(idx[c], i)
	}
	for _, arr := range idx {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return
}

// func uniqueLetterString(s string) int {
// 	var ret int
// 	var length = len(s)
// 	var result = make(map[string]int, length*(length+1)/2+1)
// 	var countUniqueChars func(string) int
// 	countUniqueChars = func(sub string) (ans int) {
// 		character := make(map[rune]int, 26)
// 		for _, v := range sub {
// 			character[v]++
// 		}
// 		for _, v := range character {
// 			if v == 1 {
// 				ans++
// 			}
// 		}

// 		return ans
// 	}
// 	for i := 0; i < len(s); i++ {
// 		for j := len(s) - 1; j >= i; j-- {
// 			result[i:j] = 0
// 		}
// 	}
// 	for _, v := range result {
// 		ret += v
// 	}
// 	return ret
// }

// @lc code=end
