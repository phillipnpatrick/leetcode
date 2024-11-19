package Solutions

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: target = 7, nums = [2,3,1,2,4,3]
			// Output: 2
			name: "test01",
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			// Input: target = 4, nums = [1,4,4]
			// Output: 1
			name: "test02",
			args: args{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			want: 1,
		},
		{
			// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
			// Output: 0
			name: "test03",
			args: args{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
