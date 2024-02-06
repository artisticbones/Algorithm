package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isValid(t *testing.T) {
	Convey("Test_isValid...", t, func() {
		So(isValid("()"), ShouldBeTrue)
		So(isValid("(]"), ShouldBeFalse)
	})
}
