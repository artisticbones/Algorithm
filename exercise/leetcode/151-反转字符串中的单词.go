package leetcode

func reverseWords(s string) string {
	// 统一处理所有字符串，保证最后一个单词能够复制。
	s = " " + s
	str := make([]rune, 0, len(s))
	for _, c := range s {
		str = append(str, c)
	}
	i, j := len(str)-1, len(str)

	res := make([]rune, 0, len(str))
	for i >= 0 {
		if str[i] == ' ' {
			if i+1 < j {
				res = append(res, str[i+1:j]...)
				res = append(res, ' ')
			}
			j = i
		}
		i--
	}
	res = res[0 : len(res)-1]
	return string(res)
}
