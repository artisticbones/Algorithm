package leetcode

func replaceSpace(s string) string {
	if s == "" {
		return ""
	}

	res := make([]rune, 0, len(s))
	for _, c := range s {
		if c == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, c)
		}
	}

	return string(res)
}
