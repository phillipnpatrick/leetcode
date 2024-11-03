package Solutions

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
			// Output: [1,2,2,3,5,6]
			name: "test01",
			args: args{
				nums1: []int {1,2,3,0,0,0},
				m: 3,
				nums2: []int {2,5,6},
				n: 3,
			},
			want: []int {1,2,2,3,5,6},
		},
		{
			// Input: nums1 = [1], m = 1, nums2 = [], n = 0
			// Output: [1]
			name: "test02",
			args: args{
				nums1: []int {1},
				m: 1,
				nums2: []int {},
				n: 0,
			},
			want: []int {1},
		},
		{
			// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
			// Output: [1]
			name: "test03",
			args: args{
				nums1: []int {0},
				m: 0,
				nums2: []int {1},
				n: 1,
			},
			want: []int {1},
		},
		{
			// Input: nums1 = [5,6,7,0,0,0], m = 3, nums2 = [1,2,3], n = 3
			// Output: [1,2,3,5,6,7]
			name: "test04",
			args: args{
				nums1: []int {5,6,7,0,0,0},
				m: 3,
				nums2: []int {1,2,3},
				n: 3,
			},
			want: []int {1,2,3,5,6,7},
		},
		{
			name: "test05",
			args: args{
				nums1: []int {4,5,6,0,0,0},
				m: 3,
				nums2: []int {1,2,3},
				n: 3,
			},
			want: []int {1,2,3,4,5,6},
		},
		{
			name: "test06",
			args: args{
				nums1: []int {1,2,3,5,6,0,0,0,0},
				m: 5,
				nums2: []int {1,2,3,4},
				n: 4,
			},
			want: []int {1,1,2,2,3,3,4,5,6},
		},
		{
			name: "test07",
			args: args{
				nums1: []int {1,2,3,0,0,0},
				m: 3,
				nums2: []int {4,5,6},
				n: 3,
			},
			want: []int {1,2,3,4,5,6},
		},
		{
			name: "test08",
			args: args{
				nums1: []int {0,0,0,0,0},
				m: 0,
				nums2: []int {1,2,3,4,5},
				n: 5,
			},
			want: []int {1,2,3,4,5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
