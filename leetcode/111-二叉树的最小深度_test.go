package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_minDepth(t *testing.T) {
	convey.Convey("Test_minDepth...", t, func() {
		convey.Convey("test case 1", func() {
			root := &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			}
			convey.So(minDepth(root), convey.ShouldEqual, 2)
		})
		convey.Convey("test case 2", func() {
			root := &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			}
			convey.So(minDepth(root), convey.ShouldEqual, 2)
		})
	})
}
