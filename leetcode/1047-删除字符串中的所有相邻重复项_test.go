package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_removeDuplicates_1047(t *testing.T) {
	Convey("Test_removeDuplicates_1047...", t, func() {
		So(removeDuplicates_1047("abbaca"), ShouldEqual, "ca")
	})
}
