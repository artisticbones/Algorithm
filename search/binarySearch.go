package search

// BinarySearch
// num: the number needs to locate
// array: The array needs to be kept in order
func BinarySearch(key int, array []int) (pos int) {
	if array == nil {
		pos = -1
		return
	}
	lo, hi := 0, len(array)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < array[mid] {
			hi = mid - 1
		} else if key > array[mid] {
			lo = mid + 1
		} else {
			lo = mid
			hi = mid
			pos = mid
		}
	}
	return
}
