package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_preorderTraversal_2(t *testing.T) {
	Convey("Test_preorderTraversal_2...", t, func() {
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
			So(preorderTraversal_2(root), ShouldResemble, []int{1, 2, 3})
		})
	})
}
