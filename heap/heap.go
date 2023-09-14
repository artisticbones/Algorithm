package heap

import "fmt"

type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Up(i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || h.less(parent, i) {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

// Push 指针传递的意义，保证函数外也起到作用
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.Up(len(*h) - 1)
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1 // 左孩子
		r := 2*i + 2 // 右孩子
		if l >= len(h) {
			break // i 已经是叶子结点，退出操作
		}
		j := 1
		if r < len(h) && h.less(r, l) {
			j = r // 右孩子为最小子节点
		}

		if h.less(i, j) {
			break // 如果父节点比叶子结点小，则不交换
		}
		h.swap(i, j)
		i = j
	}
}

func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}

	n := len(*h) - 1
	h.swap(i, n) // 用最后的元素值替换被删除的元素
	x := (*h)[n]
	*h = (*h)[0:n]

	if i == 0 || (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else {
		h.Up(i)
	}
	return x, true
}

func (h *Heap) Pop() int {
	x, ok := h.Remove(0)
	if !ok {
		return 0
	}
	return x
}

func BuildHeap(arr []int) Heap {
	h := Heap(arr)
	n := len(h)

	// 从第一个非叶子节点，到根节点
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
	return h
}

func HeapSort(array []int) {
	heap := BuildHeap(array)
	var sortedArr []int
	for len(heap) > 0 {
		sortedArr = append(sortedArr, heap.Pop())
	}
	fmt.Println(sortedArr)
}
