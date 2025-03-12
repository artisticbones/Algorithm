package leetcode

import "container/list"

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		res   = make([][]int, 0)
		queue = list.New()
		temp  = make([][]int, 0)
	)
	queue.PushBack(root)
	for queue.Len() > 0 {
		var (
			num   = queue.Len()
			level = make([]int, 0)
		)
		for i := 0; i < num; i++ {
			no := queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, no.Val)
			if no.Left != nil {
				queue.PushBack(no.Left)
			}
			if no.Right != nil {
				queue.PushBack(no.Right)
			}
		}
		temp = append(temp, level)
	}
	for i := len(temp) - 1; i >= 0; i-- {
		res = append(res, temp[i])
	}
	return res
}
