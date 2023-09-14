package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_intersection(t *testing.T) {
	Convey("Test_intersection...", t, func() {
		nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
		So(intersection(nums1, nums2), ShouldResemble, []int{2})

		nums1, nums2 = []int{4, 9, 5}, []int{9, 4, 9, 8, 4}
		So(intersection(nums1, nums2), ShouldResemble, []int{9, 4})
	})
}
