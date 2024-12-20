package Solutions

import "testing"

func Test_longestOnes(t *testing.T) {
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
			// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
			// Output: 6
			name: "test01",
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    2,
			},
			want: 6,
		},
		{
			// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
			// Output: 10
			name: "test02",
			args: args{
				nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				k:    3,
			},
			want: 10,
		},
		{
			name: "test03",
			args: args{
				nums: []int{0, 0, 0, 1},
				k:    4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
