package leetcode

import (
	"container/heap"
)

type tag struct {
	num int
	cnt int
}

type tags []tag

func (t tags) Less(i, j int) bool {
	return t[i].cnt < t[j].cnt
}

func (t tags) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tags) Len() int {
	return len(t)
}

func (t *tags) Push(x any) {
	*t = append(*t, x.(tag))
}

func (t *tags) Pop() any {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	var (
		cnt = make(map[int]int)
		res = make([]int, k)
		h   = &tags{}
	)
	for _, num := range nums {
		cnt[num]++
	}
	heap.Init(h)

	for num, c := range cnt {
		heap.Push(h, tag{
			num: num,
			cnt: c,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).(tag).num
	}

	return res
}
