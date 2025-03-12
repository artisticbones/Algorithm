package leetcode

// 左括号
var pairs = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var ret bool
	if len(s) < 1 || len(s) > 10000 {
		ret = false
		return ret
	}
	if len(s)%2 != 0 {
		ret = false
		return ret
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 { // 判断是右括号，原理在于只有右括号有键值且大于0
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				// 如果只有右括号，且栈中为空，则置false
				// 如果栈顶元素与当前元素的值不匹配，则置false
				ret = false
				return ret
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	ret = len(stack) == 0
	return ret
}
