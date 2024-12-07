/**
  @author: crane
  @since: 2024/12/2
  @desc: 包含剑指offer的所有类型定义
**/

package offer

type BinaryTreeNode struct {
	MNValue  int
	MPParent *BinaryTreeNode
	MPLeft   *BinaryTreeNode
	MPRight  *BinaryTreeNode
}

type Stack struct {
	data []int
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return 0
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

type Queue struct {
	data []int
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Push(value int) {
	q.data = append(q.data, value)
}

func (q *Queue) Pop() int {
	if len(q.data) == 0 {
		return 0
	}
	top := q.data[0]
	q.data = q.data[1:]
	return top
}
