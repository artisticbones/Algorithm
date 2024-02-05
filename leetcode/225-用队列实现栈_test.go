package leetcode

import (
	"fmt"
	"testing"
)

func TestMyStack_1(t *testing.T) {
	// ["MyStack", "push", "push", "top", "pop", "empty"]
	myS := myStack()
	fmt.Println(myS)

	myS.Push(1)
	fmt.Println(myS)

	myS.Push(2)
	fmt.Println(myS)

	x := myS.Top()
	fmt.Println("Top method: ", x)

	x = myS.Pop()
	fmt.Println("Pop method: ", x)
	fmt.Println(myS)

	empty := myS.Empty()
	fmt.Println(empty)
}
