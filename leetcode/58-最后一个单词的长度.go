package leetcode

func lengthOfLastWord(s string) int {
	var end = len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return 0
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return end - start
}

//func test01() {
//	s := "Hello World"
//	fmt.Println(lengthOfLastWord(s))
//}
