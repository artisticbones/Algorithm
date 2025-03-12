package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	Convey("Test_postorderTraversal...", t, func() {
		Convey("test case 1", func() {
			root := &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			}
			So(postorderTraversal(root), ShouldResemble, []int{3, 2, 1})
		})
	})
}
