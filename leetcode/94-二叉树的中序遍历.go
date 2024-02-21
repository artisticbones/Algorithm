package leetcode

func inorderTraversal(root *TreeNode) []int {
	var (
		temp = root
		res  = make([]int, 0)
		stk  = make([]*TreeNode, 0)
	)
	for (temp != nil) || len(stk) > 0 {
		if temp != nil {
			stk = append(stk, temp)
			temp = temp.Left
		} else {
			no := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			res = append(res, no.Val)
			temp = no.Right
		}
	}
	return res
}

func inorderTraversal_2(root *TreeNode) []int {
	var (
		res = make([]int, 0)
	)
	if root == nil {
		return nil
	}
	res = append(res, inorderTraversal_2(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal_2(root.Right)...)
	return res
}
