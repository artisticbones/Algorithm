package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_countNodes_1(t *testing.T) {
	Convey("Test_countNodes_1...", t, func() {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 6},
				Right: nil,
			},
		}
		So(countNodes_1(root), ShouldEqual, 6)
	})
}
