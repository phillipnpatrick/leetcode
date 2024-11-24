package Solutions

import "testing"

func Test_findMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: nums = [0,1]
			// Output: 2
			name: "test01",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			// Input: nums = [0,1,0]
			// Output: 2
			name: "test02",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
		{
			name: "test03",
			args: args{
				nums: []int{0, 1, 0, 1},
			},
			want: 4,
		},
		{
			name: "test04",
			args: args{
				nums: []int{0,1,1,0,1,1,1,0},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
