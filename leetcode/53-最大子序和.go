package leetcode

// method_one_Dynamic_programming
func maxSubArray(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > sum {
			sum = nums[i]
		}
	}
	return sum
}

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// method_two_Divide_and_conquer
//func maxSubArray(nums []int) int {
//	sum := 0
//	left, right := 0, len(nums) - 1
//	mid := (right - left) / 2 + left
//	for left <= right {
//		tmp := maxSubArray(nums[:mid]) + maxSubArray(nums[mid:])
//		if tmp > sum {
//			sum = tmp
//		}
//		left ++
//		right --
//	}
//	return sum
//}
