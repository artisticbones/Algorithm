package _117

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var (
		queue = list.New()
	)
	queue.PushBack(root)
	for queue.Len() > 0 {
		var (
			num        = queue.Len()
			levelNodes = make([]*Node, 0)
		)
		for i := 0; i < num; i++ {
			no := queue.Remove(queue.Front()).(*Node)
			levelNodes = append(levelNodes, no)
			if no.Left != nil {
				queue.PushBack(no.Left)
			}
			if no.Right != nil {
				queue.PushBack(no.Right)
			}
		}
		for i := 0; i < len(levelNodes)-1; i++ {
			levelNodes[i].Next = levelNodes[i+1]
		}
	}
	return root
}
