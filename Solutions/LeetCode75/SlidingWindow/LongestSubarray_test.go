package Solutions

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [1,1,0,1]
			// Output: 3
			name: "test01",
			args: args{
				nums: []int{1, 1, 0, 1},
			},
			want: 3,
		},
		{
			// Input: nums = [0,1,1,1,0,1,1,0,1]
			// Output: 5
			name: "test02",
			args: args{
				nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			},
			want: 5,
		},
		{
			// Input: nums = [1,1,1]
			// Output: 2
			name: "test03",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
