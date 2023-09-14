package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reverseList(t *testing.T) {
	Convey("Test_reverseList...", t, func() {
		Convey("test example 1", func() {
			head := &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			}
			ret := []int{5, 4, 3, 2, 1}
			p := reverseList(head)
			tmp := make([]int, 0, 1)
			for p != nil {
				tmp = append(tmp, p.Val)
				p = p.Next
			}
			So(tmp, ShouldResemble, ret)
		})
	})
}
