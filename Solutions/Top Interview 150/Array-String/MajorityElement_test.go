package Solutions

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [3,2,3]
			// Output: 3
			name: "test01",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			// Input: nums = [2,2,1,1,1,2,2]
			// Output: 2
			name: "test02",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
