package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	Convey("Test_canConstruct...", t, func() {
		So(canConstruct("a", "b"), ShouldBeFalse)
		So(canConstruct("aa", "ab"), ShouldBeFalse)
	})
}
