package leetcode

// 下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）
func nextNumber(num int) (int, int) {
	// 利用动态规划方法：
	bigger, smaller := num, num
	bigger = bigger << 1
	smaller = smaller >> 1
	return bigger, smaller
}
