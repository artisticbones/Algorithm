package leetcode

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
