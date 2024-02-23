package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	Convey("Test_averageOfLevels...", t, func() {
		Convey("test case 1", func() {
			root := &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			}
			So(averageOfLevels(root), ShouldResemble, []float64{3, 14.5, 11})
		})
	})
}
