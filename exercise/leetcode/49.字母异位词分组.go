/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
package leetcode

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0, 8)
	divide := make(map[[26]int][]string, 16)
	for _, str := range strs {
		cnt := [26]int{}
		for _, s := range str {
			cnt[s-'a']++
		}
		divide[cnt] = append(divide[cnt], str)
	}
	for _, v := range divide {
		ret = append(ret, v)
	}
	return ret
}

// @lc code=end
