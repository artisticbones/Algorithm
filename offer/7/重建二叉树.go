/**
  @author: crane
  @since: 2024/11/27
  @desc: 剑指offer 第 7 题: 重建二叉树，根据前序遍历和中序遍历重建二叉树，假设输入的前序遍历和中序遍历的结果中都不包含重复的数字。
**/

package _7

import "github.com/artisticbones/Algorithm/offer"

func construct(preorder []int, inorder []int, length int) *offer.BinaryTreeNode {
	if preorder == nil || inorder == nil || length <= 0 {
		return nil
	}
	return constructCore(0, length-1, 0, length-1, preorder, inorder)
}

// constructCore
//
//	@Description: 重建二叉树的核心算法
//	@param startPreorder int
//	@param endPreorder int
//	@param startInorder int
//	@param endInorder int
//	@param preorder []int
//	@param inorder []int
//	@return *offer.BinaryTreeNode
func constructCore(startPreorder, endPreorder, startInorder, endInorder int, preorder, inorder []int) *offer.BinaryTreeNode {
	// 前序遍历的第一个数字是根节点的值
	rootValue := preorder[startPreorder]
	root := &offer.BinaryTreeNode{MNValue: rootValue}
	if startPreorder == endPreorder {
		if startInorder == endInorder && preorder[startPreorder] == preorder[startInorder] {
			return root
		}
	}
	// 在中序遍历中找到根节点的值
	rootInorder := startInorder
	for rootInorder <= endInorder && inorder[rootInorder] != rootValue {
		rootInorder++
	}
	if rootInorder == endInorder && inorder[rootInorder] != rootValue {
		panic("Invalid input")
	}
	leftLength := rootInorder - startInorder
	leftPreorderEnd := startPreorder + leftLength
	if leftLength > 0 {
		root.MPLeft = constructCore(startPreorder+1, leftPreorderEnd, startInorder, rootInorder-1, preorder, inorder)
	}
	if leftLength < endPreorder-startPreorder {
		root.MPRight = constructCore(leftPreorderEnd+1, endPreorder, rootInorder+1, endInorder, preorder, inorder)
	}
	return root

}
