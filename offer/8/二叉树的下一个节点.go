/**
  @author: crane
  @since: 2024/12/3
  @desc: 剑指offer 第 8 题: 二叉树的下一个节点
**/

package _8

import "github.com/artisticbones/Algorithm/offer"

// GetNext
//
//	@Description: 给定一颗二叉树和其中的一个节点，如何找出中序遍历序列的下一个节点？
//	@param pNode *offer.BinaryTreeNode
//	@return *offer.BinaryTreeNode
func GetNext(pNode *offer.BinaryTreeNode) *offer.BinaryTreeNode {
	if pNode == nil {
		return nil
	}
	var pNext *offer.BinaryTreeNode
	// 如果当前节点有右子树，则下一个节点是右子树的最左节点
	if pNode.MPRight != nil {
		pRight := pNode.MPRight
		for pRight.MPLeft != nil {
			pRight = pRight.MPLeft
		}
		pNext = pRight
		// 如果当前节点没有右子树，且当前节点是父节点的左子节点，则下一个节点是父节点
	} else if pNode.MPParent != nil {
		pCurrent := pNode
		pParent := pNode.MPParent
		// 如果当前节点是父节点的右子节点，则沿着父节点向上遍历，直到找到一个节点是其父节点的左子节点
		for pParent != nil && pCurrent == pParent.MPRight {
			pCurrent = pParent
			pParent = pParent.MPParent
		}
		pNext = pParent
	}
	return pNext

}
