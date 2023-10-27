package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	Convey("Test_repeatedSubstringPattern...", t, func() {
		So(repeatedSubstringPattern("abcabc"), ShouldBeTrue)
		So(repeatedSubstringPattern("ac"), ShouldBeFalse)
		So(repeatedSubstringPattern("aa"), ShouldBeTrue)
	})
}
