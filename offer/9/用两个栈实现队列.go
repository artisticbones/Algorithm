/**
  @author: crane
  @since: 2024/12/4
  @desc: 剑指offer 第 9 题: 用两个栈实现队列
**/

package _9

import "github.com/artisticbones/Algorithm/offer"

type queue struct {
	input  *offer.Stack
	output *offer.Stack
}

// appendTail
//
//	@Description: 在队列尾部插入节点
//	@receiver q
//	@param value int
func (q *queue) appendTail(value int) {
	q.input.Push(value)
}

// deleteHead
//
//	@Description: 删除队列头部节点
//	@receiver q
func (q *queue) deleteHead() {
	if q.output.Empty() {
		for !q.input.Empty() {
			q.output.Push(q.input.Pop())
		}
	}
	if q.output.Empty() {
		panic("queue is empty")
	}
	q.output.Pop()
}
