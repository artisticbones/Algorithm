package leetcode

import (
	"container/list"
	"math"
)

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		queue = list.New()
		res   = make([]int, 0)
	)
	queue.PushBack(root)
	for queue.Len() > 0 {
		var (
			num = queue.Len()
			tag = math.MinInt
		)
		for i := 0; i < num; i++ {
			no := queue.Remove(queue.Front()).(*TreeNode)
			if tag < no.Val {
				tag = no.Val
			}
			if no.Left != nil {
				queue.PushBack(no.Left)
			}
			if no.Right != nil {
				queue.PushBack(no.Right)
			}
		}
		res = append(res, tag)
	}
	return res
}
