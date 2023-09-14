package leetcode

func min(i, j int) bool { return i < j }

func maxArea(height []int) int {
	var (
		area = 0
		i    = 0
		j    = len(height) - 1
	)

	for i <= j {
		x := j - i
		y := 0
		if min(height[i], height[j]) {
			y = height[i]
			i++
		} else {
			y = height[j]
			j--
		}
		t := x * y
		if min(area, t) {
			area = t
		}
	}

	return area
}
