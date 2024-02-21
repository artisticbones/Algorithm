package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	Convey("Test_rightSideView...", t, func() {
		Convey("test case 1", func() {
			root := &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: &TreeNode{Val: 4},
				},
			}
			So(rightSideView(root), ShouldResemble, []int{1, 3, 4})
		})
	})
}
