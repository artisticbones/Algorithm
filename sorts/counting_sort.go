package chapter_eight

/**
 * @artisticbones
 * @summary 计数排序
 * @param in slice int, 需要排序的数组
 * @param k int， 数组中最大的元素
 * @out out slice int, 存放排序完成的数组
 * @time complexity, O(k+n)
 */
func CountingSort(in []int, k int) (out []int) {
	counting := make([]int, k+1)
	out = make([]int, len(in))
	for i := range counting {
		counting[i] = 0
	}
	for i := range in {
		counting[in[i]]++
	}
	for i := 1; i < len(counting); i++ {
		counting[i] += counting[i-1]
	}
	for i := len(in) - 1; i >= 0; i-- {
		out[counting[in[i]]-1] = in[i]
		counting[in[i]]--
	}
	return
}
