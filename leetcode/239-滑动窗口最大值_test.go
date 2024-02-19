package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	Convey("Test_maxSlidingWindow", t, func() {
		Convey("test case 1", func() {
			So(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), ShouldResemble, []int{3, 3, 5, 5, 6, 7})
		})
		Convey("test case 2", func() {
			So(maxSlidingWindow([]int{1}, 1), ShouldResemble, []int{1})
		})
	})
}
