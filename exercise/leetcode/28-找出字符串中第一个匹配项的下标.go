package leetcode

func cmpStr(a, b string) bool {
	return a == b
}

// strStr 滑动窗口
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	i, j := 0, len(needle)
	for j <= len(haystack) {
		if cmpStr(haystack[i:j], needle) {
			return i
		}
		i++
		j = i + len(needle)
	}
	return -1
}
