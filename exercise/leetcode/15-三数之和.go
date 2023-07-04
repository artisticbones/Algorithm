package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 || nums == nil {
		return nil
	}
	sort.Ints(nums)
	ret := make([][]int, 0, 1)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return ret
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			t := nums[i] + nums[left] + nums[right]
			switch {
			case t == 0:
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				//去重，因为 i 不变，当此时 l取的数的值与前一个数相同，所以不用在计算，直接跳
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				//去重，因为 i不变，当此时 r 取的数的值与前一个相同，所以不用在计算
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			case t < 0:
				left++
			case t > 0:
				right--
			}
		}
	}
	return ret
}
