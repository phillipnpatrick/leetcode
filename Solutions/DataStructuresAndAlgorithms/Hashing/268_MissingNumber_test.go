package Solutions

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [3,0,1]
			// Output: 2
			name: "test01",
			args: args{
				nums: []int{3, 0, 1},
			},
			want: 2,
		},
		{
			// Input: nums = [0,1]
			// Output: 2
			name: "test02",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			// Input: nums = [9,6,4,2,3,5,7,0,1]
			// Output: 8
			name: "test03",
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
