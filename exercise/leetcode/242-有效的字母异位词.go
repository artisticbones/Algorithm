package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	anagram := make(map[rune]int, len(s))
	for _, char := range s {
		anagram[char]++
	}
	for _, char := range t {
		if _, ok := anagram[char]; !ok {
			return false
		}
		anagram[char]--
		if anagram[char] < 0 {
			return false
		}
	}
	return true
}
