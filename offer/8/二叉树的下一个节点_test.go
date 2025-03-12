/**
  @author: crane
  @since: 2024/12/4
  @desc: test cases
**/

package _8

import (
	"github.com/artisticbones/Algorithm/offer"
	"reflect"
	"testing"
)

func TestGetNext(t *testing.T) {
	type args struct {
		pNode *offer.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
		want *offer.BinaryTreeNode
	}{
		{
			name: "case 1",
			args: args{
				pNode: &offer.BinaryTreeNode{
					MNValue: 'a',
					MPLeft: &offer.BinaryTreeNode{
						MNValue: 'b',
						MPParent: &offer.BinaryTreeNode{
							MNValue: 'a',
						},
						MPLeft: &offer.BinaryTreeNode{
							MNValue: 'd',
							MPParent: &offer.BinaryTreeNode{
								MNValue: 'b',
							},
						},
						MPRight: &offer.BinaryTreeNode{
							MNValue: 'e',
							MPParent: &offer.BinaryTreeNode{
								MNValue: 'b',
							},
							MPLeft: &offer.BinaryTreeNode{
								MNValue: 'h',
								MPParent: &offer.BinaryTreeNode{
									MNValue: 'e',
								},
							},
							MPRight: &offer.BinaryTreeNode{
								MNValue: 'i',
								MPParent: &offer.BinaryTreeNode{
									MNValue: 'e',
								},
							},
						},
					},
					MPRight: &offer.BinaryTreeNode{
						MNValue: 'c',
						MPParent: &offer.BinaryTreeNode{
							MNValue: 'a',
						},
						MPLeft: &offer.BinaryTreeNode{
							MNValue: 'f',
							MPParent: &offer.BinaryTreeNode{
								MNValue: 'c',
							},
						},
						MPRight: &offer.BinaryTreeNode{
							MNValue: 'g',
							MPParent: &offer.BinaryTreeNode{
								MNValue: 'c',
							},
						},
					},
				},
			},
			want: &offer.BinaryTreeNode{
				MNValue: 'c',
				MPParent: &offer.BinaryTreeNode{
					MNValue: 'a',
				},
				MPLeft: &offer.BinaryTreeNode{
					MNValue: 'f',
					MPParent: &offer.BinaryTreeNode{
						MNValue: 'c',
					},
				},
				MPRight: &offer.BinaryTreeNode{
					MNValue: 'g',
					MPParent: &offer.BinaryTreeNode{
						MNValue: 'c',
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNext(tt.args.pNode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
