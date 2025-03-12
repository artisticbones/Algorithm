package leetcode

func searchInsert(nums []int, target int) int {
	index := len(nums)
	// 暴力查找算法时间复杂度为O(n)，故可以采用二分查找
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}
