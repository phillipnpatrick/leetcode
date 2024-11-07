package Solutions

import (
	"reflect"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			// Input: nums1 = [1,2,3], nums2 = [2,4,6]
			// Output: [[1,3],[4,6]]
			name: "test01",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4, 6},
			},
			want: [][]int{{1, 3}, {4, 6}},
		},
		{
			// Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
			// Output: [[3],[]]
			name: "test02",
			args: args{
				nums1: []int{1, 2, 3, 3},
				nums2: []int{1, 1, 2, 2},
			},
			want: [][]int{{3}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifference(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}