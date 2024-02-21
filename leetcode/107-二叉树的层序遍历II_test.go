package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	Convey("Test_levelOrderBottom...", t, func() {
		Convey("test case 1", func() {
			root := &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			}
			So(levelOrderBottom(root), ShouldResemble, [][]int{{15, 7}, {9, 20}, {3}})
		})
	})
}
