package Solutions

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			// Input: nums = [[3,1,2,4,5],[1,2,3,4],[3,4,5,6]]
			// Output: [3,4]
			name: "test01",
			args: args{
				nums: [][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}},
			},
			want: []int{3, 4},
		},
		{
			// Input: nums = [[1,2,3],[4,5,6]]
			// Output: []
			name: "test02",
			args: args{
				nums: [][]int{{1, 2, 3}, {4, 5, 6}},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
