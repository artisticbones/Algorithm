package chapter_three

func swap(i, j int) { i, j = j, i }

func Partition(array []int, p, q int) int {
	sential := array[p]
	i := p
	for j := p + 1; j < q; j++ {
		if array[j] <= sential {
			continue
		}
		i++
		swap(array[i], array[j])
	}
	swap(sential, array[i])
	return i
}

func Quicksort(array []int, p, q int) {
	if p < q {
		r := Partition(array, p, q)
		Quicksort(array, p, r-1)
		Quicksort(array, r+1, q)
	}
}
