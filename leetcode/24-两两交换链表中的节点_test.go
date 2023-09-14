package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	Convey("Test_swapPairs...", t, func() {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		}

		tmp := &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		}

		ret := swapPairs(head)
		So(ret, ShouldResemble, tmp)
	})
}
