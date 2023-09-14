package chapter_eight

func left(root int) int { return 2*root + 1 }

func right(root int) int { return 2*root + 2 }

func MaxHeapify(array []int, index int) {
	// get children's index
	l := left(index)
	r := right(index)
	largest := 0
	if l < len(array) && array[l] > array[index] {
		largest = l
	} else {
		largest = index
	}
	if r < len(array) && array[r] > array[largest] {
		largest = r
	}
	if largest != index {
		swap(&array[index], &array[largest])
		MaxHeapify(array, largest)
	}
}

func BuildMaxHeap(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		MaxHeapify(array, i)
	}
}

func Heapsort(array []int) {
	BuildMaxHeap(array)
	for i := len(array) - 1; i > 0; i-- {
		swap(&array[0], &array[i])
		MaxHeapify(array, 0)
	}
}
