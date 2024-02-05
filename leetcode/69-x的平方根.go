package leetcode

import (
	_ "math"
)

// method_one: 袖珍计算器算法是一种用指数函数exp和对数函数ln代替平方根函数的方法。
// 我们通过有限的可以使用的数学函数，得到我们想要计算的结果。
//
//	func mySqrt(x int) int {
//		var ret int
//		if x == 0 {
//			ret = 0
//			return ret
//		}
//		ret = int(math.Exp(0.5 * math.Log(float64(x))))
//		if (ret+1)*(ret+1) <= x {
//			return ret + 1
//		}
//		return ret
//	}
//
// method_ two: binary search
// 由于x平方根的整数部分ans是满足k的平方 <= x的最大k值，因此我们可以对k进行二分查找，从而得到答案。
// 二分查找的下界为0，上界可以粗略的设定为x。在二分查找的每一步中，我们只需要比较中间元素 mid 的平方与 x 的大小关系，
// 并通过比较的结果调整上下界的范围。由于是整数运算，不会存在误差，因此在得到最终的答案 ans 后，不需要尝试ans + 1
func mySqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
