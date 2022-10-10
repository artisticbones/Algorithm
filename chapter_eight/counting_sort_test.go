package chapter_eight

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	type args struct {
		in []int
		k  int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				in: []int{1, 3, 5, 9, 7, 4},
				k:  9,
			},
			wantOut: []int{1, 3, 4, 5, 7, 9},
		},
		{
			name: "test2",
			args: args{
				in: []int{1, 3, 3, 9, 9, 4},
				k:  9,
			},
			wantOut: []int{1, 3, 3, 4, 9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := CountingSort(tt.args.in, tt.args.k); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("CountingSort() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
