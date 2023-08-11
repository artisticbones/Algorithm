package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isHappy(t *testing.T) {
	Convey("Test_isHappy...", t, func() {
		So(isHappy(19), ShouldBeTrue)
		So(isHappy(2), ShouldBeFalse)
	})
}
