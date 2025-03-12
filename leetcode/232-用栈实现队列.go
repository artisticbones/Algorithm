package leetcode

type MyQueue struct {
	in  []int
	out []int
}

func constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	x := this.Peek()
	this.out = this.out[:len(this.out)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	if len(this.out) != 0 {
		return this.out[len(this.out)-1]
	}
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	this.in = make([]int, 0)

	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
