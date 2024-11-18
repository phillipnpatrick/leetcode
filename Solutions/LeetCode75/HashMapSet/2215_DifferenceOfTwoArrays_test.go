package Solutions

import (
	"reflect"
	"sort"
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
		{
			name: "test03",
			args: args{
				nums1: []int{-3, 6, -5, 4, 5, 5},
				nums2: []int{6, 6, -3, -3, 3, 5},
			},
			want: [][]int{{-5, 4}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Ints(tt.want[0])
			sort.Ints(tt.want[1])
			got := findDifference(tt.args.nums1, tt.args.nums2);
			sort.Ints(got[0])
			sort.Ints(got[1])
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
