package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	type fields struct {
		m      int
		n      int
		result int
	}
	Convey("Test_uniquePaths...", t, func() {
		tests := []fields{
			{m: 1, n: 1, result: 1},
			{m: 3, n: 7, result: 28},
		}
		for _, test := range tests {
			So(uniquePaths(test.m, test.n), ShouldEqual, test.result)
		}
	})
}
