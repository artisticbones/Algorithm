package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	convey.Convey("Test_isSymmetric...", t, func() {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		}
		convey.So(isSymmetric(root), convey.ShouldBeTrue)
	})
}
