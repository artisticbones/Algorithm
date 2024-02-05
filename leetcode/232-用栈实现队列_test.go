package leetcode

import (
	"fmt"
	"testing"
)

func TestMyQueue_1(t *testing.T) {
	myQ := constructor()
	fmt.Println(myQ)

	myQ.Push(1)
	fmt.Println(myQ)

	myQ.Push(2)
	fmt.Println(myQ)

	x := myQ.Peek()
	fmt.Println("Peek method: ", x)
	fmt.Println(myQ)

	x = myQ.Pop()
	fmt.Println("Pop method: ", x)
	fmt.Println(myQ)

	empty := myQ.Empty()
	fmt.Println(empty)
}

func TestMyQueue_2(t *testing.T) {
	myQ := constructor()
	fmt.Println(myQ)

	myQ.Push(1)
	fmt.Println(myQ)

	x := myQ.Pop()
	fmt.Println("Pop method: ", x)
	fmt.Println(myQ)

	empty := myQ.Empty()
	fmt.Println(empty)
}

func TestMyQueue_3(t *testing.T) {
	// ["MyQueue","push","push","push","push","pop","push","pop","pop","pop","pop"]
	// [[],[1],[2],[3],[4],[],[5],[],[],[],[]]
	myQ := constructor()
	fmt.Println(myQ)

	myQ.Push(1)
	fmt.Println(myQ)

	myQ.Push(2)
	fmt.Println(myQ)

	myQ.Push(3)
	fmt.Println(myQ)

	myQ.Push(4)
	fmt.Println(myQ)

	x := myQ.Pop()
	fmt.Println("Pop method: ", x)
	fmt.Println(myQ)

	myQ.Push(5)
	fmt.Println(myQ)

	x = myQ.Pop()
	fmt.Println("Pop method: ", x)
	fmt.Println(myQ)

	x = myQ.Pop()
	fmt.Println("Pop method: ", x)
	fmt.Println(myQ)

	x = myQ.Pop()
	fmt.Println("Pop method: ", x)
	fmt.Println(myQ)

	x = myQ.Pop()
	fmt.Println("Pop method: ", x)
	fmt.Println(myQ)
}
