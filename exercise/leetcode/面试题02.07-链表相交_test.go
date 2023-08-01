package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	Convey("Test_getIntersectionNode...", t, func() {
		Convey("test example 1...", func() {
			headA := Ints2List([]int{4, 1, 8, 4, 5})
			headB := Ints2List([]int{5, 0, 1, 8, 4, 5})
			ret := Ints2List([]int{8, 4, 5})
			So(getIntersectionNode(headA, headB), ShouldResemble, ret)
		})
	})
}
