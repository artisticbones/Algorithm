package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 0 {
		return 0
	}
	ret := make([]int, 0, len(nums1)+len(nums2))
	i, j, median := 0, 0, 0.0

	for {
		if i > len(nums1)-1 || j > len(nums2)-1 {
			break
		}
		if nums1[i] < nums2[j] {
			ret = append(ret, nums1[i])
			i++
		} else {
			ret = append(ret, nums2[j])
			j++
		}
	}

	if i == len(nums1) {
		ret = append(ret, nums2[j:]...)
	}
	if j == len(nums2) {
		ret = append(ret, nums1[i:]...)
	}

	if len(ret)%2 == 0 {
		middleIndex1 := len(ret) / 2
		middleIndex2 := middleIndex1 - 1
		median = float64(ret[middleIndex1]+ret[middleIndex2]) / 2
	} else {
		median = float64(ret[(len(ret)-1)/2])
	}
	return median
}
