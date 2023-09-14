package leetcode

func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	var i, j = 0, len(nums) - 1

	for i <= j {
		var temp = 0
		for i < j {
			if nums[i] != val {
				i++
			} else {
				if i < j {
					temp = nums[j]
					nums[j] = nums[i]
					nums[i] = temp
				}
				break
			}
		}
		for i < j {
			if nums[j] == val {
				j--
			} else {
				if j > i {
					temp = nums[j]
					nums[j] = nums[i]
					nums[i] = temp
				}
				break
			}
		}
	}
	return i
}
