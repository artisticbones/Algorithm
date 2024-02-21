package leetcode

func postorderTraversal(root *TreeNode) []int {
	var (
		res = make([]int, 0)
	)
	if root == nil {
		return nil
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

func postorderTraversal_2(root *TreeNode) []int {
	var (
		res  = make([]int, 0)
		stk  = make([]*TreeNode, 0)
		temp = root
	)

	for len(stk) > 0 || temp != nil {
		if temp != nil {
			res = append([]int{temp.Val}, res...)
			stk = append(stk, temp)
			temp = temp.Right
		} else {
			temp = stk[len(stk)-1].Left
			stk = stk[:len(stk)-1]
		}
	}

	return res
}
