package leetcode

func reverseString_2(s []rune) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseStr(s string, k int) string {
	if k == 0 {
		return s
	}
	chars := make([]rune, 0, len(s))
	for _, b := range s {
		chars = append(chars, b)
	}
	for i := 0; i <= len(chars)/k; i++ {
		if i%2 == 0 {
			if (i+1)*k < len(chars) {
				reverseString_2(chars[i*k : (i+1)*k])
			} else {
				reverseString_2(chars[i*k:])
			}
		} else {
			continue
		}
	}
	return string(chars)
}
