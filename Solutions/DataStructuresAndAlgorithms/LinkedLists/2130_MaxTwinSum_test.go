package Solutions

import "testing"

func Test_pairSum(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	// Input: head = [5,4,3,2,1]
		// 	// Output: 6
		// 	name: "test00",
		// 	args: args{head: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}}},
		// 	want: 6,
		// },
		{
			// Input: head = [5,4,2,1]
			// Output: 6
			name: "test01",
			args: args{head: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}},
			want: 6,
		},
		{
			// Input: head = [4,2,2,3]
			// Output: 7
			name: "test02",
			args: args{head: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}},
			want: 7,
		},
		{
			// Input: head = [1,100000]
			// Output: 100001
			name: "test03",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 100000, Next: nil}}},
			want: 100001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSum(tt.args.head); got != tt.want {
				t.Errorf("pairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
