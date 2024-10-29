package Solutions

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [3,2,2,3], val = 3
			// Output: 2, nums = [2,2,_,_]
			name: "test01",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			// Input: nums = [0,1,2,2,3,0,4,2], val = 2
			// Output: 5, nums = [0,1,4,0,3,_,_,_]
			name: "test02",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "test03",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  5,
			},
			want: 8,
		},
		{
			name: "test04",
			args: args{
				nums: []int{1},
				val:  1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
