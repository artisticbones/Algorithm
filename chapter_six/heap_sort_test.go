package chaptersix

import (
	"fmt"
	"testing"
)

func TestHeapsort(t *testing.T) {
	type args struct {
		array []int
	}
	// var (
	// 	a, b, c, d, e = 1, 4, 3, 5, 9
	// )
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				array: []int{1, 5, 4, 3, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Heapsort(tt.args.array)
			for _, c := range tt.args.array {
				fmt.Println(c)
			}
		})
	}
}
