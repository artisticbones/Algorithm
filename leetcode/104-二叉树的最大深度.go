package leetcode

import "container/list"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		res   = 0
		queue = list.New()
	)
	queue.PushBack(root)
	for queue.Len() > 0 {
		var (
			num = queue.Len()
		)
		for i := 0; i < num; i++ {
			no := queue.Remove(queue.Front()).(*TreeNode)
			if no.Left != nil {
				queue.PushBack(no.Left)
			}
			if no.Right != nil {
				queue.PushBack(no.Right)
			}
		}
		res++
	}
	return res
}
