package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "test success",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "test 2",
			args: args{
				nums1: []int{},
				nums2: []int{},
			},
			want: 0.0,
		},
		{
			name: "test 3",
			args: args{
				nums1: []int{1},
				nums2: []int{7, 8, 10, 11},
			},
			want: 8,
		},
		{
			name: "test 4",
			args: args{
				nums1: []int{},
				nums2: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
