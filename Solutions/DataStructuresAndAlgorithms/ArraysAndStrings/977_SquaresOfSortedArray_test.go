package Solutions

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{
				nums: []int{-4,-1,0,3,10},
			},
			want: []int{0,1,9,16,100},
		},
		{
			name: "test02",
			args: args{
				nums: []int{-7,-3,2,3,11},
			},
			want: []int{4,9,9,49,121},
		},
		{
			name: "test03",
			args: args{
				nums: []int{-4,-1,0,3,5,6,7,8,9,10},
			},
			want: []int{0,1,9,16,25,36,49,64,81,100},
		},
		{
			name: "test04",
			args: args{
				nums: []int{-5,-4,-3,-2,-1,0,1},
			},
			want: []int{0,1,1,4,9,16,25},
		},
		{
			name: "test05",
			args: args{
				nums: []int{0,2},
			},
			want: []int{0,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
