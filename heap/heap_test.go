package heap

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHeapSort(t *testing.T) {
	Convey("TestHeapSort...", t, func() {
		arr := []int{33, 24, 8, 3, 10, 15, 16, 15, 30, 17, 19}
		Sort(arr)
	})
}
