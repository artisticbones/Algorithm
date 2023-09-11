package kmp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_kmpSearch(t *testing.T) {
	Convey("Test_kmpSearch...", t, func() {
		So(kmpSearch("sadbutsad", "sad"), ShouldEqual, 0)
		So(kmpSearch("leetcode", "leeto"), ShouldEqual, -1)
		So(kmpSearch("hello", "ll"), ShouldEqual, 2)
		So(kmpSearch("abc", "c"), ShouldEqual, 2)
	})
}
