/**
  @author: crane
  @since: 2024/11/20
  @desc: //todo
**/

package _4

import "testing"

func Test_find(t *testing.T) {
	type args struct {
		matrix  [][]int
		rows    int
		columns int
		target  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{1, 2, 8, 9},
					{2, 4, 9, 12},
					{4, 7, 10, 13},
					{6, 8, 11, 15},
				},
				rows:    4,
				columns: 4,
				target:  7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find(tt.args.matrix, tt.args.rows, tt.args.columns, tt.args.target); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
