package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最长同值路径长度必定为某一节点的左最长同值有向路径长度与右最长同值有向路径长度之和。
func longestUnivaluePath(root *TreeNode) int {
	var dfs func(*TreeNode) int
	var ans int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 左孩子
		left := dfs(node.Left)
		// 右孩子
		right := dfs(node.Right)
		// 进行深度优先遍历
		leftAns, rightAns := 0, 0
		if node.Left != nil && node.Val == node.Left.Val {
			leftAns = left + 1
		}
		if node.Right != nil && node.Val == node.Right.Val {
			rightAns = right + 1
		}
		ans = max(ans, leftAns+rightAns)
		return max(leftAns, rightAns)
	}
	dfs(root)
	return ans
}
