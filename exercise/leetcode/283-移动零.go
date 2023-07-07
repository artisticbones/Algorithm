package leetcode

func moveZeroes(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	var j = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			t := nums[i]
			nums[i] = nums[j]
			nums[j] = t
			j++
		}
	}
}
