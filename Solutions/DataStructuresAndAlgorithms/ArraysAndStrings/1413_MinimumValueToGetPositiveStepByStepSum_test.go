package Solutions

import "testing"

func Test_minStartValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [-3,2,-3,4,2]
			// Output: 5
			name: "test01",
			args: args{
				nums: []int{-3, 2, -3, 4, 2},
			},
			want: 5,
		},
		{
			// Input: nums = [1,2]
			// Output: 1
			name: "test02",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		},
		{
			// Input: nums = [1,-2,-3]
			// Output: 5
			name: "test03",
			args: args{
				nums: []int{1, -2, -3},
			},
			want: 5,
		},
		{
			name: "test04",
			args: args{
				nums: []int{-12},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStartValue(tt.args.nums); got != tt.want {
				t.Errorf("minStartValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
