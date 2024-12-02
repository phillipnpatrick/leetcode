package Solutions

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			// Input: head = [1,2,6,3,4,5,6], val = 6
			// Output: [1,2,3,4,5]
			name: "test01",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}},
				val:  6,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		},
		{
			// Input: head = [], val = 1
			// Output: []
			name: "test02",
			args: args{
				head: &ListNode{Val: 1, Next: nil},
				val:  1,
			},
			want: nil,
		},
		{
			// Input: head = [7,7,7,7], val = 7
			// Output: []
			name: "test03",
			args: args{
				head: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: nil}}}},
				val:  7,
			},
			want: nil,
		},
		{
			name: "test04",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
				val:  1,
			},
			want: &ListNode{Val: 2, Next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
