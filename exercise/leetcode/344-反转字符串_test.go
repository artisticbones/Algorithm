package leetcode

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reverseString(t *testing.T) {
	Convey("Test_reverseString...", t, func() {
		s := []byte{'h', 'e', 'l', 'l', 'o'}
		fmt.Printf("%v\n", s)

		reverseString(s)
		fmt.Printf("%v\n", s)

		s = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
		fmt.Printf("%v\n", s)

		reverseString(s)
		fmt.Printf("%v\n", s)
	})
}
