package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	Convey("Test_topKFrequent...", t, func() {
		So(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), ShouldResemble, []int{1, 2})
		So(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 1), ShouldResemble, []int{1})
	})
}
