package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	Convey("Test_removeNthFromEnd...", t, func() {
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
		tmp := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		}
		result := removeNthFromEnd(head, 2)
		So(result, ShouldResemble, tmp)
	})
}
