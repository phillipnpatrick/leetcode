package Solutions

import "testing"

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			// Input: nums = [1,12,-5,-6,50,3], k = 4
			// Output: 12.75000
			name: "test01",
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75000,
		},
		{
			// Input: nums = [5], k = 1
			// Output: 5.00000
			name: "test02",
			args: args{
				nums: []int{5},
				k:    1,
			},
			want: 5.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
