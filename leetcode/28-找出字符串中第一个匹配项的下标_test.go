package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_strStr(t *testing.T) {
	convey.Convey("Test_strStr...", t, func() {
		convey.So(strStr("sadbutsad", "sad"), convey.ShouldEqual, 0)
		convey.So(strStr("leetcode", "leeto"), convey.ShouldEqual, -1)
		convey.So(strStr("hello", "ll"), convey.ShouldEqual, 2)
		convey.So(strStr("abc", "c"), convey.ShouldEqual, 2)
	})
}
