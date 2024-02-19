package leetcode

import "sort"

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
		res = make([]int, 0)
		h   = make([]tag, 0)
	)
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		h = append(h, tag{
			num: num,
			cnt: c,
		})
	}
	hp := tags(h)
	sort.Sort(hp)
	for i := 0; i < k; i++ {
		res = append(res, hp.Pop().(tag).num)
	}

	return res
}
