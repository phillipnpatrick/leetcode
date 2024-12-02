package Solutions

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			// Input: head = [1,2,3,4,5], left = 2, right = 4
			// Output: [1,4,3,2,5]
			name: "test01",
			args: args{
				head:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
				left:  2,
				right: 4,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: nil}}}}},
		},
		{
			// Input: head = [5], left = 1, right = 1
			// Output: [5]
			name: "test02",
			args: args{
				head:  &ListNode{Val: 5, Next: nil},
				left:  1,
				right: 1,
			},
			want: &ListNode{Val: 5, Next: nil},
		},
		{
			name: "test03",
			args: args{
				head:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}}}}}}}},
				left:  3,
				right: 5,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}}}}}}}},
		},
		{
			name: "test04",
			args: args{
				head:  &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}},
				left:  1,
				right: 1,
			},
			want: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}},
		},
		{
			name: "test05",
			args: args{
				head:  &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}},
				left:  1,
				right: 2,
			},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
