package Solutions

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			// Input: head = [1,2,3,4]
			// Output: [2,1,4,3]
			name: "test01",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}},
		},
		{
			// Input: head = []
			// Output: []
			name: "test02",
			args: args{head: nil},
			want: nil,
		},
		{
			// Input: head = [1]
			// Output: [1]
			name: "test03",
			args: args{head: &ListNode{Val: 1, Next: nil}},
			want: &ListNode{Val: 1, Next: nil},
		},
		{
			// Input: head = [1,2,3]
			// Output: [2,1,3]
			name: "test04",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}},
		},
		{
			// Input: head = [2,5,3,4,6,2,2]
			// Output: [5,2,4,3,2,6,2]
			name: "test05",
			args: args{head: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: nil}}}}}}}},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 2, Next: nil}}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
