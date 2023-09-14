package leetcode

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConstructor(t *testing.T) {
	Convey("Test all", t, func() {
		l := Constructor()
		l.AddAtHead(1)
		fmt.Printf("%v", l)
		l.AddAtTail(3)
		fmt.Printf("%v", l)
		l.AddAtIndex(1, 2)
		fmt.Printf("%v", l)
		l.Get(1)
		l.DeleteAtIndex(1)
		fmt.Printf("%v", l)
		l.Get(1)
	})
}
