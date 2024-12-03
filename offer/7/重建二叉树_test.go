/**
  @author: crane
  @since: 2024/12/2
  @desc: test cases
**/

package _7

import (
	"github.com/artisticbones/Algorithm/offer"
	"reflect"
	"testing"
)

func Test_construct(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
		length   int
	}
	tests := []struct {
		name string
		args args
		want *offer.BinaryTreeNode
	}{
		{
			name: "case 1",
			args: args{
				preorder: []int{1, 2, 4, 7, 3, 5, 6, 8},
				inorder:  []int{4, 7, 2, 1, 5, 3, 8, 6},
				length:   8,
			},
			want: &offer.BinaryTreeNode{
				MNValue: 1,
				MPLeft: &offer.BinaryTreeNode{
					MNValue: 2,
					MPLeft: &offer.BinaryTreeNode{
						MNValue: 4,
						MPRight: &offer.BinaryTreeNode{
							MNValue: 7,
						},
					},
				},
				MPRight: &offer.BinaryTreeNode{
					MNValue: 3,
					MPLeft: &offer.BinaryTreeNode{
						MNValue: 5,
					},
					MPRight: &offer.BinaryTreeNode{
						MNValue: 6,
						MPLeft: &offer.BinaryTreeNode{
							MNValue: 8,
						},
					},
				},
			},
		},
		{
			name: "case 2",
			args: args{
				preorder: []int{1, 2, 3, 4, 5},
				inorder:  []int{5, 4, 3, 2, 1},
				length:   5,
			},
			want: &offer.BinaryTreeNode{
				MNValue: 1,
				MPRight: &offer.BinaryTreeNode{
					MNValue: 2,
					MPRight: &offer.BinaryTreeNode{
						MNValue: 3,
						MPRight: &offer.BinaryTreeNode{
							MNValue: 4,
							MPRight: &offer.BinaryTreeNode{
								MNValue: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "case 3",
			args: args{
				preorder: []int{1},
				inorder:  []int{1},
				length:   1,
			},
			want: &offer.BinaryTreeNode{
				MNValue: 1,
			},
		},
		{
			name: "case 4",
			args: args{
				preorder: []int{1, 2, 3, 4, 5},
				inorder:  []int{1, 2, 3, 4, 5},
				length:   5,
			},
			want: &offer.BinaryTreeNode{
				MNValue: 1,
				MPLeft:  nil,
				MPRight: &offer.BinaryTreeNode{
					MNValue: 2,
					MPLeft:  nil,
					MPRight: &offer.BinaryTreeNode{
						MNValue: 3,
						MPLeft:  nil,
						MPRight: &offer.BinaryTreeNode{
							MNValue: 4,
							MPLeft:  nil,
							MPRight: &offer.BinaryTreeNode{
								MNValue: 5,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct(tt.args.preorder, tt.args.inorder, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct() = %v, want %v", got, tt.want)
			}
		})
	}
}
