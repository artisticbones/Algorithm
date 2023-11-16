package leetcode

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
*/

func binarySearch(nums []int, start, end int, target int) int {
	mid := start + (end-start)>>1
	if mid == start && mid == end {
		if nums[mid] < target {
			return start + 1
		} else {
			return start
		}
	}
	switch {
	case nums[mid] == target:
		return mid
	case nums[mid] < target:
		start = mid + 1
		return binarySearch(nums, start, end, target)
	case nums[mid] > target:
		end = mid
		return binarySearch(nums, start, end, target)
	}
	return mid
}

func searchAndInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}
