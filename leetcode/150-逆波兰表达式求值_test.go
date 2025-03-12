package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	Convey("Test_evalRPN...", t, func() {
		Convey("test case 1", func() {
			So(evalRPN([]string{"2", "1", "+", "3", "*"}), ShouldEqual, 9)
		})
		Convey("test case 2", func() {
			So(evalRPN([]string{"4", "13", "5", "/", "+"}), ShouldEqual, 6)
		})
		Convey("test case 3", func() {
			So(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}), ShouldEqual, 22)
		})
	})
}
