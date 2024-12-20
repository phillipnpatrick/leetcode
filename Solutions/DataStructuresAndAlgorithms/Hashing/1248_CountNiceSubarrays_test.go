package Solutions

import "testing"

func Test_numberOfSubarrays(t *testing.T) {
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
			// Input: nums = [1,1,2,1,1], k = 3
			// Output: 2
			name: "test01",
			args: args{
				nums: []int{1, 1, 2, 1, 1},
				k:    3,
			},
			want: 2,
		},
		{
			// Input: nums = [2,4,6], k = 1
			// Output: 0
			name: "test02",
			args: args{
				nums: []int{2, 4, 6},
				k:    1,
			},
			want: 0,
		},
		{
			// Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
			// Output: 16
			name: "test03",
			args: args{
				nums: []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
				k:    2,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numberOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
