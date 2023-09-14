package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reverseLeftWords(t *testing.T) {
	Convey("Test_reverseLeftWords...", t, func() {
		So(reverseLeftWords("abcdefg", 2), ShouldEqual, "cdefgab")
	})
}
