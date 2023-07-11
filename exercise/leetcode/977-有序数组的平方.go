package leetcode

import "math"

func pow2(num int) int { return num * num }

func sortedSquares(nums []int) []int {
	if nums == nil {
		return nil
	}
	ret := make([]int, 0, len(nums))
	index := 0
	for _, num := range nums {
		if num <= 0 {
			index++
		} else {
			break
		}
	}
	left, right := nums[:index], nums[index:]
	i, j := len(left)-1, 0
	for {
		if i < 0 || j > len(right)-1 {
			break
		}
		if math.Abs(float64(left[i])) <= math.Abs(float64(right[j])) {
			ret = append(ret, pow2(left[i]))
			i--
		} else {
			ret = append(ret, pow2(right[j]))
			j++
		}
	}
	for i >= 0 {
		ret = append(ret, pow2(left[i]))
		i--
	}
	for j <= len(right)-1 {
		ret = append(ret, pow2(right[j]))
		j++
	}
	return ret
}
