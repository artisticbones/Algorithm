package leetcode

/*
给你一个整数数组nums,请计算数组的中心下标 。
数组中心下标是数组的一个下标,其左侧所有元素相加的和等于右侧所有元素相加的和。
如果中心下标位于数组最左端,那么左侧数之和视为0,因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
*/

func pivotIndex(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	total := 0
	for _, num := range nums {
		total += num
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		if total-nums[i] == 2*left {
			return i
		}
		left += nums[i]
	}
	return -1
}
