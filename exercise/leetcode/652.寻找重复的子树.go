/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
 */
package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	type pairs struct {
		node *TreeNode
		idx  int // 使用三元组的index对子树进行唯一标识
	}
	seem := make(map[[3]int]pairs, 4)
	repeat := make(map[*TreeNode]struct{}, 1)
	idx := 0
	var dfs func(*TreeNode) int
	dfs = func(no *TreeNode) int {
		if no == nil {
			return 0
		}
		// 构造子树
		tri := [3]int{no.Val, dfs(no.Left), dfs(no.Right)}
		if pair, ok := seem[tri]; ok {
			// 子树重复，直接return不需要在向下遍历
			repeat[pair.node] = struct{}{}
			return pair.idx
		}
		idx++
		seem[tri] = pairs{node: no, idx: idx}
		return idx
	}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}

// @lc code=end
