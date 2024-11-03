package Solutions

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// Example 1:
		// Input: nums = [0,1,0,3,12]
		// Output: [1,3,12,0,0]
		{
			name: "test01",
			args: args{nums: []int{0,1,0,3,12}},
			want: []int{1,3,12,0,0},
		},
		// Input: nums = [0]
		// Output: [0]
		{
			name: "test02",
			args: args{nums: []int{0}},
			want: []int{0},
		},
		{
			name: "test03",
			args: args{nums: []int{2,4,6,0,8,0,10,12,0}},
			want: []int{2,4,6,8,10,12,0,0,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
