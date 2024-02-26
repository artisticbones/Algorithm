package _16

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_connect(t *testing.T) {
	Convey("Test_connect...", t, func() {
		root := &Node{
			Val: 1,
			Left: &Node{
				Val:   2,
				Left:  &Node{Val: 4},
				Right: &Node{Val: 5},
				Next:  nil,
			},
			Right: &Node{
				Val:   3,
				Left:  &Node{Val: 6},
				Right: &Node{Val: 7},
				Next:  nil,
			},
			Next: nil,
		}
		fmt.Println(connect(root))
	})
}
