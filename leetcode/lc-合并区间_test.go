package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeSet(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSet(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
