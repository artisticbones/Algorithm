package leetcode

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum1, sum2 := map[int]int{}, map[int]int{}

	for _, s1 := range nums1 {
		for _, s2 := range nums2 {
			sum1[s1+s2]++
		}
	}
	for _, s3 := range nums3 {
		for _, s4 := range nums4 {
			sum2[s3+s4]++
		}
	}

	total := 0
	for s1, n1 := range sum1 {
		if n2, ok := sum2[-s1]; ok {
			total += n1 * n2
		}
	}
	return total
}
