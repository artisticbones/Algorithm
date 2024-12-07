/**
  @author: crane
  @since: 2024/12/7
  @desc: 剑指offer 第 9 题: 用两个队列实现栈
**/

package _9

import "github.com/artisticbones/Algorithm/offer"

type stack struct {
	data1 *offer.Queue
	data2 *offer.Queue
}

func (s *stack) pop() int {
	if s.data1.Empty() && s.data2.Empty() {
		panic("stack is empty")
	}
	if s.data1.Empty() {
		for !s.data2.Empty() {
			s.data1.Push(s.data2.Pop())
		}
		return s.data2.Pop()
	} else {
		for !s.data1.Empty() {
			s.data2.Push(s.data1.Pop())
		}
		return s.data1.Pop()
	}
}

func (s *stack) push(value int) {
	if s.data1.Empty() {
		s.data2.Push(value)
	} else {
		s.data1.Push(value)
	}
}
