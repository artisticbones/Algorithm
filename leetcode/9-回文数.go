/**
 * 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 * 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
 */

package leetcode

func isPalindrome(x int) bool {
	ret := true
	reverseNum := 0
	// 方法一(数学法):
	// 1. 将数反转与元数据比较
	if x < 0 {
		ret = false
		return ret
	}
	tmp := x
	for tmp != 0 {
		digit := tmp % 10
		tmp /= 10
		reverseNum = reverseNum*10 + digit
	}
	if x != reverseNum {
		ret = false
		return ret
	}
	return ret
}
