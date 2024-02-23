package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_largestValues(t *testing.T) {
	Convey("Test_largestValues...", t, func() {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: &TreeNode{Val: 9},
			},
		}
		So(largestValues(root), ShouldResemble, []int{1, 3, 9})
	})
}
