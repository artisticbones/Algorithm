package leetcode

import (
	"sort"
)

/*
以数组intervals表示若干个区间的集合，其中单个区间为intervals[i] = [starti, endi].
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
*/

type ArraySort [][]int

func (a ArraySort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ArraySort) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a ArraySort) Len() int {
	return len(a)
}

func mergeSet(intervals [][]int) [][]int {
	sort.Sort(ArraySort(intervals))
	ret := make([][]int, 0, len(intervals))
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if tmp[1] >= intervals[i][0] {
			tmp[1] = max(tmp[1], intervals[i][1])
		} else {
			ret = append(ret, tmp)
			tmp = intervals[i]
		}
	}
	ret = append(ret, tmp)
	return ret
}
