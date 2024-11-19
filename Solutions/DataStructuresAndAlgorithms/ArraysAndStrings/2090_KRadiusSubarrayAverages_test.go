package Solutions

import (
	"reflect"
	"testing"
)

func Test_getAverages(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			// Input: nums = [7,4,3,9,1,8,5,2,6], k = 3
			// Output: [-1,-1,-1,5,4,4,-1,-1,-1]
			name: "test01",
			args: args{
				nums: []int{7, 4, 3, 9, 1, 8, 5, 2, 6},
				k:    3,
			},
			want: []int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
		},
		{
			// Input: nums = [100000], k = 0
			// Output: [100000]
			name: "test02",
			args: args{
				nums: []int{100000},
				k:    0,
			},
			want: []int{100000},
		},
		{
			// Input: nums = [8], k = 100000
			// Output: [-1]
			name: "test03",
			args: args{
				nums: []int{8},
				k:    100000,
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAverages(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
