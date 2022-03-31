package chapter_two

// MergeSort
// 分解： 分解待排序的 n 个元素的序列成各具 n/2 个元素的两个字序列
// 解决： 使用归并排序递归地排序两个子序列
// 合并： 合并两个已排序的子序列以产生已排序的答案
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
