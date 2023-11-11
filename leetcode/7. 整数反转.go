package leetcode

import (
	"math"
)

func reverse(x int) int {
	// 弹出末尾数字
	var ret int
	for x != 0 {
		if ret > (math.MaxInt32/10) || ret < (math.MinInt32/10) {
			ret = 0
			return ret
		}
		digit := x % 10
		x /= 10
		ret = ret*10 + digit
	}
	return ret
}
