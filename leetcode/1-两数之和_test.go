package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_twoSum(t *testing.T) {
	Convey("Test_twoSum...", t, func() {
		So(twoSum([]int{2, 7, 11, 15}, 9), ShouldResemble, []int{0, 1})
		So(twoSum([]int{3, 2, 4}, 6), ShouldResemble, []int{1, 2})
		So(twoSum([]int{3, 3}, 6), ShouldResemble, []int{0, 1})
	})
}
