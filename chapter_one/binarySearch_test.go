package chapter_one

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	Convey("TestBinarySearch...", t, func() {
		key := 1
		array := []int{0, 1, 2, 3, 4, 5}
		So(BinarySearch(key, array), ShouldEqual, 1)
	})
}
