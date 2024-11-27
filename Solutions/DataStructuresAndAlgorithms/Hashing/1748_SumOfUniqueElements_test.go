package Solutions

import "testing"

func Test_sumOfUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [1,2,3,2]
			// Output: 4
			name: "test01",
			args: args{
				nums: []int{1, 2, 3, 2},
			},
			want: 4,
		},
		{
			// Input: nums = [1,1,1,1,1]
			// Output: 0
			name: "test02",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: 0,
		},
		{
			// Input: nums = [1,2,3,4,5]
			// Output: 15
			name: "test03",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfUnique(tt.args.nums); got != tt.want {
				t.Errorf("sumOfUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
