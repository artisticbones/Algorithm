package leetcode

func minAB(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func lcp(str1 string, str2 string) string {
	length := minAB(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func longestCommonPrefix(strs []string) string {
	var ret string
	if len(strs) < 1 || len(strs) > 200 {
		return ret
	}
	ret = strs[0]
	for _, v := range strs {
		if len(v) < 0 || len(v) > 200 {
			continue
		}
		ret = lcp(ret, v)
	}
	return ret
}
