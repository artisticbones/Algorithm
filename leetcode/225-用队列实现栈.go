package leetcode

type MyStack struct {
	queue1 []int
	queue2 []int
}

func myStack() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	this.queue2 = append(this.queue2, this.queue1...)

	this.queue1 = this.queue2
	this.queue2 = make([]int, 0)
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	x := this.Top()
	if len(this.queue1) == 1 {
		this.queue1 = make([]int, 0)
	} else {
		this.queue1 = this.queue1[1:]
	}
	return x
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
