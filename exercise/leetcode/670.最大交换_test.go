/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_maximumSwap(t *testing.T) {
	Convey("Test_maximumSwap...", t, func() {
		num := 98368
		So(maximumSwap(num), ShouldEqual, 98863)
	})
}
