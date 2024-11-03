package Solutions

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [1,1,2]
			// Output: 2, nums = [1,2,_]
			name: "test01",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			// Input: nums = [0,0,1,1,1,2,2,3,3,4]
			// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
			name: "test02",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
