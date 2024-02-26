package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	Convey("Test_maxDepth...", t, func() {
		root := &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}
		So(maxDepth(root), ShouldEqual, 3)
	})
}
