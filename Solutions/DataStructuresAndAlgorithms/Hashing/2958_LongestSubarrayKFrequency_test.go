package Solutions

import "testing"

func Test_maxSubarrayLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [1,2,3,1,2,3,1,2], k = 2
			// Output: 6
			name: "test01",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3, 1, 2},
				k:    2,
			},
			want: 6,
		},
		{
			// Input: nums = [1,2,1,2,1,2,1,2], k = 1
			// Output: 2
			name: "test02",
			args: args{
				nums: []int{1, 2, 1, 2, 1, 2, 1, 2},
				k:    1,
			},
			want: 2,
		},
		{
			// Input: nums = [5,5,5,5,5,5,5], k = 4
			// Output: 4
			name: "test03",
			args: args{
				nums: []int{5, 5, 5, 5, 5, 5, 5},
				k:    4,
			},
			want: 4,
		},
		{
			name: "test04",
			args: args{
				nums: []int{1, 4, 4, 3},
				k:    1,
			},
			want: 2,
		},
		{
			name: "test05",
			args: args{
				nums: []int{2, 2, 3},
				k:    1,
			},
			want: 2,
		},
		{
			name: "test06",
			args: args{
				nums: []int{1, 1, 1, 3},
				k:    2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarrayLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
