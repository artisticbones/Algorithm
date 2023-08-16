package leetcode

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_fourSum(t *testing.T) {
	Convey("Test_fourSum...", t, func() {
		tests := []struct {
			nums   []int
			target int
		}{
			{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
		}
		for _, test := range tests {
			fmt.Printf("%v\n", fourSum(test.nums, test.target))
		}
	})
}
