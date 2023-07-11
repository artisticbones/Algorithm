package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test success",
			args: args{nums: []int{-4, -1, 0, 3, 10}},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "test left",
			args: args{nums: []int{-5, -3, -2, -1}},
			want: []int{1, 4, 9, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
