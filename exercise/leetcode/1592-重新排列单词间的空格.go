/*
 * @lc app=leetcode.cn id=1592 lang=golang
 *
 * [1592] 重新排列单词间的空格
 */
package leetcode

import (
	"strings"
)

// @lc code=start
func reorderSpaces(text string) string {
	words := strings.Fields(text)
	blankNum := strings.Count(text, " ")
	lw := len(words) - 1
	if lw == 0 {
		return words[0] + strings.Repeat(" ", blankNum)
	}
	return strings.Join(words, strings.Repeat(" ", blankNum/lw)) + strings.Repeat(" ", blankNum%lw)
}

// @lc code=end
