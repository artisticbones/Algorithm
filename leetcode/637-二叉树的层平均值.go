package leetcode

import "container/list"

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var (
		res   = make([]float64, 0)
		queue = list.New()
	)
	queue.PushBack(root)
	for queue.Len() > 0 {
		var (
			num = queue.Len()
			sum = 0.0
		)
		for i := 0; i < num; i++ {
			no := queue.Remove(queue.Front()).(*TreeNode)
			sum += float64(no.Val)
			if no.Left != nil {
				queue.PushBack(no.Left)
			}
			if no.Right != nil {
				queue.PushBack(no.Right)
			}
		}
		res = append(res, sum/float64(num))
	}
	return res
}
