package leetcode

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

type node struct {
	val  int
	next *node
}

type MyLinkedList struct {
	length int
	head   *node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		length: 0,
		head:   nil,
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.length {
		return -1
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.length, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.length {
		return
	}
	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	p := &node{
		val:  val,
		next: prev.next,
	}
	prev.next = p
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.length {
		return
	}
	p := l.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	p.next, l.length = p.next.next, l.length-1
}
