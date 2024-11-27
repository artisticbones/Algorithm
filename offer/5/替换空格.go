/**
  @author: crane
  @since: 2024/11/26
  @desc: 剑指offer 第 5 题: 替换空格
**/

package _5

// replaceBlank
//
//	@Description: 替换空格，先统计空格数量，然后从后往前替换
//	@param s string
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
