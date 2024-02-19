package leetcode

type deque struct {
	nums []int
}

func newDeque() *deque {
	return &deque{nums: make([]int, 0)}
}

func (d *deque) Front() int {
	return d.nums[0]
}

func (d *deque) Tail() int {
	return d.nums[len(d.nums)-1]
}

func (d *deque) Empty() bool {
	return len(d.nums) == 0
}

func (d *deque) Pop(val int) {
	if !d.Empty() && val == d.Front() {
		d.nums = d.nums[1:]
	}
}

func (d *deque) Push(x int) {
	for !d.Empty() && x > d.Tail() {
		d.nums = d.nums[:len(d.nums)-1]
	}
	d.nums = append(d.nums, x)
}

func maxSlidingWindow(nums []int, k int) []int {
	var (
		res   = make([]int, 0)
		queue = newDeque()
	)

	// 先将前k个元素入队
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个值的最大值
	res = append(res, queue.Front())
	for i := k; i < len(nums); i++ {
		// 每次循环删除窗口最左侧元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加窗口最右侧元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}

	return res
}
