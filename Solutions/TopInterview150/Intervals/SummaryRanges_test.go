package Solutions

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			// Example 1:
			// Input: nums = [0,1,2,4,5,7]
			// Output: ["0->2","4->5","7"]
			name: "test01",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 7},
			},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			// Example 2:
			// Input: nums = [0,2,3,4,6,8,9]
			// Output: ["0","2->4","6","8->9"]
			name: "test02",
			args: args{
				nums: []int{0, 2, 3, 4, 6, 8, 9},
			},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "test03",
			args: args{
				nums: []int{-1},
			},
			want: []string{"-1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
