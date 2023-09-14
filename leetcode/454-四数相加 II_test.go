package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_fourSumCount(t *testing.T) {
	Convey("Test_fourSumCount...", t, func() {
		So(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}), ShouldEqual, 2)
		So(fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}), ShouldEqual, 1)
		So(fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}), ShouldEqual, 6)
	})
}
