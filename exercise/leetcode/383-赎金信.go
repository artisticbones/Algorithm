package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	chars := make(map[rune]int, len(magazine))
	for _, r := range magazine {
		chars[r]++
	}
	for _, r := range ransomNote {
		if _, ok := chars[r]; !ok {
			return false
		} else {
			chars[r]--
			if chars[r] < 0 {
				return false
			}
		}
	}
	return true
}
