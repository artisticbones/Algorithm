package leetcode

import (
	"fmt"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		level = 0
	)
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		fmt.Println(1<<level + 1)
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		fmt.Println(bits)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
