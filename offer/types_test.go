/**
  @author: crane
  @since: 2024/12/7
  @desc: test cases
**/

package offer

import "testing"

func TestQueue_Pop(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			fields: fields{
				data: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				data: tt.fields.data,
			}
			if got := q.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
