package queue

/*
使用含有尾部指针的链表来实现:
	enqueue(value) —— 在尾部添加值
	dequeue() —— 删除最早添加的元素并返回其值（首部元素）
	empty()
使用固定大小的数组实现：
	enqueue(value) —— 在可容的情况下添加元素到尾部
	dequeue() —— 删除最早添加的元素并返回其值
	empty()
	full()
花销：
	在糟糕的实现情况下，使用链表所实现的队列，其入列和出列的时间复杂度将会是 O(n)。因为，你需要找到下一个元素，以致循环整个队列
	enqueue：O(1)（平摊（amortized）、链表和数组 [探测（probing）]）
	dequeue：O(1)（链表和数组）
	empty：O(1)（链表和数组）
*/

type Node struct {
	data int
	next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	size int
}

func (q *Queue) Enqueue(value int) {
	data := &Node{
		data: value,
		next: nil,
	}
	if q == nil {
		q.Head = data
		q.Tail = data
		q.size++
		return
	}
	q.Tail.next = data
	q.Tail = data
	q.size++
}

func (q *Queue) Dequeue() int {
	var data *Node
	if q.Head != nil {
		data = q.Head
	}
	q.Head = q.Head.next
	q.size--
	return data.data
}

func (q *Queue) IsEmpty() bool {
	if q.Head == nil || q.size == 0 {
		return true
	}
	return false
}

type ArrayQueue struct {
	head  int
	tail  int
	items [16]int
}

func (a *ArrayQueue) Size() int {
	return len(a.items)
}

func (a *ArrayQueue) Enqueue(value int) {
	if (a.tail%a.Size())+1 == a.head%a.Size() {
		panic("out of range")
	}
	a.items[a.tail] = value
	a.tail++
}

func (a *ArrayQueue) Dequeue() int {
	if a.head%a.Size() == a.tail%a.Size() {
		panic("queue is empty")
	}
	data := a.items[a.head]
	a.head++
	return data
}

func (a *ArrayQueue) IsEmpty() bool {
	if a.head%a.Size() == a.tail%a.Size() {
		return true
	}
	return false
}
