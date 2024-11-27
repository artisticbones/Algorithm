/**
  @author: crane
  @since: 2024/11/26
  @desc: test cases
**/

package _5

import "testing"

func Test_replaceBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{s: "We are happy."},
			want: "We%20are%20happy.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceBlank(tt.args.s); got != tt.want {
				t.Errorf("replaceBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}
