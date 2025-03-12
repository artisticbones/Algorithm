package leetcode

import "container/list"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		queue = list.New()
		res   = 0
	)
	queue.PushBack(root)
	for queue.Len() > 0 {
		num := queue.Len()
		res++
		for i := 0; i < num; i++ {
			no := queue.Remove(queue.Front()).(*TreeNode)
			if no.Left == nil && no.Right == nil {
				return res
			}
			if no.Left != nil {
				queue.PushBack(no.Left)
			}
			if no.Right != nil {
				queue.PushBack(no.Right)
			}
		}
	}
	return res
}
