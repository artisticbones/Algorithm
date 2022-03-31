package chapter_two

func MergeSort(arr []int) []int {
	if arr == nil {
		return nil
	}
	if len(arr) < 2 {
		return arr
	}

	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (ret []int) {
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		ret = append(ret, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		ret = append(ret, right[0])
		right = right[1:]
	}
	return
}
