package leetcode

import "strconv"

/*
 */
func addBinary(a string, b string) string {
	// 从低位开始运算，模拟列加法
	// carry 表示上一个位置的进位
	// 则每一位的答案为(carry + a[i] + b[i]) % 2
	// 下一位的进位为(carry + a[i] + b[i]) / 2
	// 为了让各个位置对齐，你可以先反转这个代表二进制数字的字符串，然后低下标对应低位，高下标对应高位。
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := maxAB(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func maxAB(x, y int) int {
	if x > y {
		return x
	}
	return y
}
