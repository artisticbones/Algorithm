package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	Convey("Test_reverseWords_1...", t, func() {
		res := reverseWords("  hello world  ")
		So(len(res), ShouldEqual, 11)
		So(res, ShouldEqual, "world hello")
	})
	Convey("Test_reverseWords_2...", t, func() {
		res := reverseWords("the sky is blue")
		//So(len(res), ShouldEqual, 13)
		So(res, ShouldEqual, "blue is sky the")
	})
}
