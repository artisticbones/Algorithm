package tree

import (
	"container/list"
	"fmt"
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

// GetNodeCount 查找树上的节点数
func (t *tree) GetNodeCount() int {
	total := 0
	return total
}

// PrintValues 从小到大打印树中节点的值
func (t *tree) PrintValues() {

}

func (t *tree) DeleteTree() {

}

// IsInTree 判断值是否存在于树中
func (t *tree) IsInTree(val int) bool {
	return false
}

// GetHeight 返回节点所在的高度（如果只有一个节点，那么高度则为1）
func (t *tree) GetHeight() int {
	height := 0
	return height
}

// GetMin 返回树上的最小值
func (t *tree) GetMin() int {

}

// GetMax 返回树上的最大值
func (t *tree) GetMax() int {

}

func (t *tree) IsBinarySearchTree() bool {
	return false
}

func (t *tree) DeleteValue(val int) {

}

// GetSuccessor 返回给定值的后继者，若没有则返回-1
func (t *tree) GetSuccessor(val int) int {

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
