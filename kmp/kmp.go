package kmp

func generateNext(pattern string) (next []int) {
	j, next := 0, make([]int, len(pattern)+1)
	next[0] = 0

	for i := 1; i < len(pattern); i++ { // 注意 i 要从1开始，不要与j重合
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1] // 注意这里，是要找前一位的对应的回退位置了
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return
}

func kmpSearch(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	j, next := 0, generateNext(needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
