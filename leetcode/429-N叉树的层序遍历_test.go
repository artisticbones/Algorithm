package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_levelOrder_429(t *testing.T) {
	Convey("Test_levelOrder_429", t, func() {
		root := &Node{
			Val: 1,
			Children: []*Node{
				{
					Val: 3,
					Children: []*Node{
						{Val: 5},
						{Val: 6},
					},
				},
				{Val: 2},
				{Val: 4},
			},
		}
		So(levelOrder_429(root), ShouldResemble, [][]int{{1}, {3, 2, 4}, {5, 6}})
	})
}
