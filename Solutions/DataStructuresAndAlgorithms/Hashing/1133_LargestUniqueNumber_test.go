package Solutions

import "testing"

func Test_largestUniqueNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [5,7,3,9,4,9,8,3,1]
			// Output: 8
			name: "test01",
			args: args{
				nums: []int{5, 7, 3, 9, 4, 9, 8, 3, 1},
			},
			want: 8,
		},
		{
			// Input: nums = [9,9,8,8]
			// Output: -1
			name: "test02",
			args: args{
				nums: []int{9, 9, 8, 8},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestUniqueNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestUniqueNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
