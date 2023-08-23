package list

/*
size() —— 返回链表中数据元素的个数
empty() —— 若链表为空则返回一个布尔值 true
value_at(index) —— 返回第 n 个元素的值（从0开始计算）
push_front(value) —— 添加元素到链表的首部
pop_front() —— 删除首部元素并返回其值
push_back(value) —— 添加元素到链表的尾部
pop_back() —— 删除尾部元素并返回其值
front() —— 返回首部元素的值
back() —— 返回尾部元素的值
insert(index, value) —— 插入值到指定的索引，并把当前索引的元素指向到新的元素
erase(index) —— 删除指定索引的节点
value_n_from_end(n) —— 返回倒数第 n 个节点的值
reverse() —— 逆序链表
remove_value(value) —— 删除链表中指定值的第一个元素
*/

type Node struct {
	Data int
	Next *Node
}

type LinkList struct {
	size int
	Head *Node
}

func (l *LinkList) Size() int {
	return l.size
}

func (l *LinkList) Empty() bool {
	return l.size == 0
}

func (l *LinkList) ValueAt(index int) int {
	if l == nil {
		panic("the pointer you access is nil")
	}
	if index < 0 || index > (l.size-1) {
		panic("out of range")
	}
	p := l.Head
	for index != 0 {
		p = p.Next
		index--
	}
	return p.Data
}

func (l *LinkList) PushFront(value int) {
	p := &Node{
		Data: value,
	}
	if l == nil {
		l.Head = p
	} else {
		p.Next = l.Head.Next
		l.Head = p
	}
	l.size++
}

func (l *LinkList) PopFront() int {
	if l == nil {
		panic("the pointer you access is nil")
	}
	if l.size == 0 || l.size == 1 {
		panic("cannot pop any element from the linkList")
	}
	data := l.Head.Data
	l.Head = l.Head.Next
	return data
}

func (l *LinkList) PushBack(value int) {
	p := &Node{
		Data: value,
	}
	if l == nil {
		l.Head = p
	} else {
		q := l.Head
		for q.Next != nil {
			q = q.Next
		}
		q.Next = p
	}
	l.size++
}

func (l *LinkList) PopBack() int {
	if l == nil {
		panic("the pointer you access is nil")
	}
	if l.size == 0 || l.size == 1 {
		panic("cannot pop any element from the linkList")
	}
	p := l.Head
	for p.Next.Next != nil {
		p = p.Next
	}
	data := p.Data
	l.size--
	p.Next = nil
	return data
}

func (l *LinkList) Front() int {
	if l == nil {
		panic("the pointer you access is nil")
	}
	return l.Head.Data
}

func (l *LinkList) Back() int {
	if l == nil {
		panic("the pointer you access is nil")
	}
	p := l.Head
	for p.Next != nil {
		p = p.Next
	}
	return p.Data
}

func (l *LinkList) Insert(index, value int) {
	if l == nil {
		panic("the pointer you access is nil")
	}
	if index < 0 || index > l.size-1 {
		panic("index out of range")
	}
	p := &Node{
		Data: value,
		Next: nil,
	}
	q := l.Head
	for index > 0 {
		q = q.Next
		index--
	}
	p.Next = q.Next
	q.Next = p
	l.size++
}

func (l *LinkList) Erase(index int) {
	if l == nil {
		panic("the pointer you access is nil")
	}
	if index < 0 || index > l.size-1 {
		panic("index out of range")
	}
	q := l.Head
	for index > 0 {
		q = q.Next
		index--
	}
	q.Next = q.Next.Next
	l.size--
}

func (l *LinkList) ValueNFromEnd(n int) int {
	if l == nil {
		panic("the pointer you access is nil")
	}
	index := l.size - n - 1
	if index < 0 || index > l.size-1 {
		panic("index out of range")
	}
	q := l.Head
	for index != 0 {
		q = q.Next
		index--
	}
	return q.Data
}

func (l *LinkList) Reverse() {

}

func (l *LinkList) RemoveValue(value int) bool {
	if l == nil {
		panic("the pointer you access is nil")
	}
	p := l.Head
	for p.Next.Next != nil {
		if p.Data == value {
			p.Next = p.Next.Next
			l.size--
			return true
		} else {
			continue
		}
	}
	if p.Next.Data == value {
		p.Next = nil
		l.size--
		return true
	}
	return false
}
