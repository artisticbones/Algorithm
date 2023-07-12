package leetcode

import "math"

func minX(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen(target int, nums []int) int {
	if nums == nil {
		return 0
	}
	ans, start, end, sum := math.MaxInt32, 0, 0, 0
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			ans = minX(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
