package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		res   = make([][]int, 0)
		queue = make([]*TreeNode, 0)
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		var (
			level = make([]int, 0)
			num   = len(queue)
		)
		for i := 0; i < num; i++ {
			no := queue[0]
			if num <= 1 {
				queue = make([]*TreeNode, 0)
			} else {
				queue = queue[1:]
			}
			level = append(level, no.Val)
			if no.Left != nil {
				queue = append(queue, no.Left)
			}
			if no.Right != nil {
				queue = append(queue, no.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
