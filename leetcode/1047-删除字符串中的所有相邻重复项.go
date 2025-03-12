package leetcode

func removeDuplicates_1047(s string) string {
	var stack = make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] != stack[len(stack)-1] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
