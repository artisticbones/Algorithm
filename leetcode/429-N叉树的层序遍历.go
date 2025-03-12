package leetcode

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder_429(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var (
		res   = make([][]int, 0)
		queue = list.New()
	)
	queue.PushBack(root)
	for queue.Len() > 0 {
		var (
			num   = queue.Len()
			level = make([]int, 0)
		)
		for i := 0; i < num; i++ {
			no := queue.Remove(queue.Front()).(*Node)
			level = append(level, no.Val)
			for j := 0; j < len(no.Children); j++ {
				queue.PushBack(no.Children[j])
			}
		}
		res = append(res, level)
	}
	return res
}
