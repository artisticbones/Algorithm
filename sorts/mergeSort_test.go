package chapter_eight

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	Convey("TestMergeSort...", t, func() {
		start := []int{0, 5, 3, 1, 2, 4}
		end := []int{0, 1, 2, 3, 4, 5}

		So(MergeSort(start), ShouldResemble, end)
		log.Println(MergeSort(start))
	})
}
