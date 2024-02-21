package leetcode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res   = make([]int, 0)
		queue = make([]*TreeNode, 0)
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		var (
			num = len(queue)
		)
		for i := 0; i < num; i++ {
			no := queue[0]
			if num <= 1 {
				queue = make([]*TreeNode, 0)
			} else {
				queue = queue[1:]
			}
			if no.Left != nil {
				queue = append(queue, no.Left)
			}
			if no.Right != nil {
				queue = append(queue, no.Right)
			}
			if i == num-1 {
				res = append(res, no.Val)
			}
		}
	}

	return res
}
