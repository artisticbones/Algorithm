package tree

import (
	"container/list"
	"fmt"
	"math"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func NewBSTNode(val int) *treeNode {
	return &treeNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}

type tree struct {
	root *treeNode
}

func NewBinarySearchTree() *tree {
	return &tree{}
}

func (t *tree) insertRecursive(root *treeNode, val int) *treeNode {
	if root == nil {
		return &treeNode{val: val, right: nil, left: nil}
	}
	if val <= root.val {
		root.right = t.insertRecursive(root.right, val)
	} else {
		root.left = t.insertRecursive(root.left, val)
	}
	return root
}

func (t *tree) Insert_2(val int) {
	t.root = t.insertRecursive(t.root, val)
}

func (t *tree) Insert(val int) {
	if t.root == nil {
		t.root = &treeNode{
			val:   val,
			left:  nil,
			right: nil,
		}
	}
	cur := t.root
	for {
		if val <= cur.val {
			if cur.left != nil {
				cur = cur.left
			} else {
				cur.left = &treeNode{val: val, left: nil, right: nil}
				return
			}
		} else {
			if cur.right != nil {
				cur = cur.right
			} else {
				cur.right = &treeNode{val: val, left: nil, right: nil}
				return
			}
		}
	}
}

func (t *tree) getNodeCount(root *treeNode) int {
	if root == nil {
		return 0
	}
	return 1 + t.getNodeCount(root.right) + t.getNodeCount(root.left)
}

// GetNodeCount 查找树上的节点数
func (t *tree) GetNodeCount() int {
	return t.getNodeCount(t.root)
}

// PrintValues 从小到大打印树中节点的值
func (t *tree) PrintValues() {
	t.dfsInOrder(t.root)
}

func (t *tree) DeleteTree() {
	t.root = nil
}

func (t *tree) search(node *treeNode, val int) *treeNode {
	if node == nil {
		return node
	}
	switch {
	case val == node.val:
		return node
	case val < node.val:
		return t.search(node.left, val)
	case val > node.val:
		return t.search(node.right, val)
	}
	return nil
}

// IsInTree 判断值是否存在于树中
func (t *tree) IsInTree(val int) bool {
	return t.search(t.root, val) != nil
}

func (t *tree) getHeight(root *treeNode) int {
	if root == nil {
		return -1
	}
	return int(math.Max(float64(t.getHeight(root.left)), float64(t.getHeight(root.right)))) + 1
}

// GetHeight 返回节点所在的高度（如果只有一个节点，那么高度则为1）
// height of node: 从节点到叶子结点的最长简单路径边的条数
// depth of node: 从节点到根节点的最长简单路径边的条数
func (t *tree) GetHeight() int {
	height := 0
	return height
}

func (t *tree) getMinRecursive(root *treeNode) int {
	if root == nil {
		return -1
	}
	if root.left == nil {
		return root.val
	}
	return t.getMinRecursive(root.left)
}

func (t *tree) getMinIterate(cur *treeNode) int {
	if cur == nil {
		return -1
	}
	for cur.left != nil {
		cur = cur.left
	}
	return cur.val
}

func (t *tree) findMin(node *treeNode) *treeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

// GetMin 返回树上的最小值
func (t *tree) GetMin() int {
	if t.root == nil {
		return -1
	}
	return t.findMin(t.root).val
}

func (t *tree) findMax(node *treeNode) *treeNode {
	for node.right != nil {
		node = node.right
	}
	return node
}

// GetMax 返回树上的最大值
func (t *tree) GetMax() int {
	if t.root == nil {
		return -1
	}
	return t.findMax(t.root).val
}

func (t *tree) isBstUtil(root *treeNode, minValue, maxValue int) bool {
	if root == nil {
		return true
	}
	if root.val < minValue && root.val > maxValue {
		return false
	}
	return t.isBstUtil(root.left, minValue, root.val) && t.isBstUtil(root.right, root.val, maxValue)
}

func (t *tree) IsBinarySearchTree() bool {
	return t.isBstUtil(t.root, -math.MaxInt64, math.MaxInt64)
}

func (t *tree) deleteValue(root *treeNode, val int) *treeNode {
	if root == nil {
		return root
	}
	if val < root.val {
		root.left = t.deleteValue(root.left, val)
	} else if val > root.val {
		root.right = t.deleteValue(root.right, val)
	} else {
		//Get ready to be deleted
		if root.left == nil && root.right == nil {
			// case 1: the node has no child
			root = nil
		} else if root.left == nil {
			// case 2: the node has one child
			root = root.right
		} else if root.right == nil {
			// case 2: the node has one child
			root = root.left
		} else {
			// case 3: the node has three child
			minNode := t.findMin(root.right)
			root.val = minNode.val
			root.right = t.deleteValue(root.right, minNode.val)
		}
	}
	return root
}

func (t *tree) DeleteValue(val int) {
	t.deleteValue(t.root, val)
}

// getSuccessor find in order successor in a BST
func (t *tree) getSuccessor(root *treeNode, val int) *treeNode {
	// Search the Node - O(H)
	current := t.search(root, val)
	if current == nil {
		return nil
	}
	if current.right != nil {
		// Case 1: the node has right tree
		return t.findMin(current.right) // find the deepest left node in the right subtree
	} else {
		// Case 2: No right subtree
		var (
			ancestor            = root
			successor *treeNode = nil
		)
		for ancestor != current {
			if current.val < ancestor.val {
				successor = ancestor // so far this is the deepest node for which current node is in left
				ancestor = ancestor.left
			} else {
				ancestor = ancestor.right
			}
		}
		return successor
	}
}

// GetSuccessor 返回给定值的后继者，若没有则返回-1
func (t *tree) GetSuccessor(val int) int {
	return t.getSuccessor(t.root, val).val
}

// BFS 广度优先遍历
func (t *tree) BFS() {
	if t == nil {
		return
	}

	queue := list.New()
	queue.PushBack(t)

	for queue.Len() > 0 {
		n := queue.Remove(queue.Front()).(*treeNode)
		fmt.Printf("%d\n", n.val)

		if n.left != nil {
			queue.PushBack(n.left)
		}
		if n.right != nil {
			queue.PushBack(n.right)
		}
	}
}

func (t *tree) dfsPreOrder(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d\n", node.val)
	t.dfsPreOrder(node.left)
	t.dfsPreOrder(node.right)
}

// DFSPreOrder 深度优先遍历的先根序
func (t *tree) DFSPreOrder() {
	t.dfsPreOrder(t.root)
}

func (t *tree) dfsInOrder(node *treeNode) {
	if node == nil {
		return
	}

	t.dfsInOrder(node.left)
	fmt.Printf("%d\n", node.val)
	t.dfsInOrder(node.right)
}

// DFSInOrder 深度优先遍历的中根序
func (t *tree) DFSInOrder() {
	t.dfsInOrder(t.root)
}

func (t *tree) dfsPostOrder(node *treeNode) {
	if node == nil {
		return
	}
	t.dfsPostOrder(node.left)
	t.dfsPostOrder(node.right)
	fmt.Printf("%d\n", node.val)
}

// DFSPostOrder 深度优先遍历的后根序
func (t *tree) DFSPostOrder() {
	t.dfsPostOrder(t.root)
}
