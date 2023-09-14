package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	symbols := map[int]int{}
	ret := make([]int, 0, 8)
	for _, i := range nums1 {
		symbols[i] = 1
	}
	for _, i := range nums2 {
		if symbols[i] == 0 {
			continue
		} else {
			ret = append(ret, i)
			symbols[i] = 0
		}
	}
	return ret
}
