package heap

import (
	"sort"
)

type Interface interface {
	sort.Interface
	Insert(x any)    // add x as element Len()
	ExtractMax() any // remove and return element 0.
	GetMax() any     //	return the maximum value but not delete it
}

func Heapify(x Interface) {
	n := x.Len()
	for i := n/2 - 1; i >= 0; i-- {
		shiftDown(x, i, n)
	}
}

func Insert(x Interface, n any) {
	x.Insert(n)
	shiftUp(x, x.Len()-1)
}

func ExtractMax(x Interface) any {
	n := x.Len() - 1
	x.Swap(0, n)
	shiftDown(x, 0, n)
	return x.ExtractMax()
}

func GetMax(x Interface) any {
	return x.GetMax()
}

func GetSize(x Interface) int {
	return x.Len()
}

func IsEmpty(x Interface) bool {
	return x.Len() == 0
}

func Remove(x Interface, index int) any {
	n := x.Len() - 1
	if n != index {
		x.Swap(index, n)
		if !shiftDown(x, index, n) {
			shiftUp(x, index)
		}
	}
	return x.ExtractMax()
}

func shiftUp(x Interface, index int) {
	for {
		// according to complete tree, conduct the parent node index
		parent := (index - 1) / 2
		// if index equals to parent, then the node is root
		// if node for the index is bigger than parent's node, then against the heap rule
		if index == parent || !x.Less(index, parent) {
			break
		}
		x.Swap(index, parent)

		index = parent
	}
}

func shiftDown(x Interface, index, n int) bool {
	i0 := index // store index in order to compare with ending i0
	for {
		child1 := 2*i0 + 1             // left child
		if child1 >= n || child1 < 0 { // prevent the case that child1 int overflow
			break
		}

		temp := child1
		if child2 := child1 + 1; child2 < n && x.Less(child2, child1) {
			temp = child2
		}

		if !x.Less(temp, i0) {
			break
		}
		x.Swap(i0, temp)
		i0 = temp
	}
	return i0 > index
}
