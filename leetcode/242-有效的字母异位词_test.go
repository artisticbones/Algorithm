package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	Convey("Test_isAnagram...", t, func() {
		So(isAnagram("anagram", "nagaram"), ShouldBeTrue)
		So(isAnagram("rat", "car"), ShouldBeFalse)
	})
}
