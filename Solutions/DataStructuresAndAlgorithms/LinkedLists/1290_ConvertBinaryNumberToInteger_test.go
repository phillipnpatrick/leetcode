package Solutions

import "testing"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: head = [1,0,1]
			// Output: 5
			name: "test01",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}},
			},
			want: 5,
		},
		{
			// Input: head = [0]
			// Output: 0
			name: "test02",
			args: args{
				head: &ListNode{Val: 0, Next: nil},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
