package Solutions

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: nums = [1,2,3,1]
			// Output: true
			name: "test01",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			// Input: nums = [1,2,3,4]
			// Output: false
			name: "test02",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			// Input: nums = [1,1,1,3,3,4,3,2,4,2]
			// Output: true
			name: "test02",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
