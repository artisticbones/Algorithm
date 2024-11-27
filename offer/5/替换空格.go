/**
  @author: crane
  @since: 2024/11/26
  @desc: 剑指offer 第 5 题: 替换空格
**/

package _5

// replaceBlank
//
//		@Description: 替换空格，先统计空格数量，然后从后往前替换，在合并两个数组时（包含字符串），如果从前往后替换，会导致多次移动元素，时间复杂度为O(n^2)
//	 则我们可以考虑从后往前替换，这样只需要移动一次元素，时间复杂度为O(n)
//		@param s string
func replaceBlank(s string) string {
	if s == "" {
		return ""
	}
	oldLen := len(s)
	blankCount := 0
	for _, c := range s {
		if c == ' ' {
			blankCount++
		}
	}
	newLen := oldLen + blankCount*2
	newStr := make([]byte, newLen)
	oldIndex, newIndex := oldLen-1, newLen-1
	for oldIndex >= 0 && newIndex >= oldIndex {
		if s[oldIndex] == ' ' {
			newStr[newIndex] = '0'
			newStr[newIndex-1] = '2'
			newStr[newIndex-2] = '%'
			newIndex -= 3
		} else {
			newStr[newIndex] = s[oldIndex]
			newIndex--
		}
		oldIndex--
	}
	return string(newStr)
}
