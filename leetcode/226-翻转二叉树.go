package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left, right := invertTree(root.Left), invertTree(root.Right)
	root.Left, root.Right = right, left

	return root
}
