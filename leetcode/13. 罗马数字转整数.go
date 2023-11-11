package leetcode

var byteToN = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var ret = 0
	n := len(s)
	for k := range s {
		value := byteToN[s[k]]
		if (k < n-1) && (value < byteToN[s[k+1]]) {
			ret -= value
		} else {
			ret += value
		}
	}
	return ret
}
