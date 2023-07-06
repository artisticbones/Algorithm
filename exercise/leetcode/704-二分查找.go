package leetcode

func search(nums []int, target int) int {
	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}
	start, end := 0, len(nums)-1
	for {
		mid := start + (end-start)>>1
		if mid == start && mid == end {
			if nums[mid] != target {
				return -1
			} else {
				return mid
			}
		}
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			start = mid + 1
			continue
		case nums[mid] > target:
			end = mid
			continue
		}
	}
}
