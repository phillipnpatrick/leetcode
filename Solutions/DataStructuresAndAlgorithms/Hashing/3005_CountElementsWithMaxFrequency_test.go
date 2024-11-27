package Solutions

import "testing"

func Test_maxFrequencyElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [1,2,2,3,1,4]
			// Output: 4
			name: "test01",
			args: args{
				nums: []int{1, 2, 2, 3, 1, 4},
			},
			want: 4,
		},
		{
			// Input: nums = [1,2,3,4,5]
			// Output: 5
			name: "test02",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequencyElements(tt.args.nums); got != tt.want {
				t.Errorf("maxFrequencyElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
