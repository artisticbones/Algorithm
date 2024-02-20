package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var (
		res = make([]int, 0)
	)
	if root == nil {
		return nil
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

type stack []any

func (s *stack) push(x any) {
	*s = append(*s, x)
}

func (s *stack) pop() any {
	if s.empty() {
		return nil
	}
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func preorderTraversal_2(root *TreeNode) []int {
	var (
		res = make([]int, 0)
	)
	if root == nil {
		return nil
	}
	s := &stack{}
	s.push(root)
	for !s.empty() {
		node := s.pop().(*TreeNode)
		res = append(res, node.Val)
		if node.Right != nil {
			s.push(node.Right)
		}
		if node.Left != nil {
			s.push(node.Left)
		}
	}
	return res
}
