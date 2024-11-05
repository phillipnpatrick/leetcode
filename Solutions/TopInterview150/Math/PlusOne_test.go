package Solutions

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			// Input: digits = [1,2,3]
			// Output: [1,2,4]
			name: "test01",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			// Input: digits = [4,3,2,1]
			// Output: [4,3,2,2]
			name: "test02",
			args: args{
				digits: []int{4, 3, 2, 1},
			},
			want: []int{4, 3, 2, 2},
		},
		{
			// Input: digits = [9]
			// Output: [1,0]
			name: "test03",
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
		{
			name: "test04",
			args: args{
				digits: []int{9,8,9,9},
			},
			want: []int{9,9,0,0},
		},
		{
			name: "test05",
			args: args{
				digits: []int{9,9,9,9},
			},
			want: []int{1,0,0,0,0},
		},
		{
			name: "test06",
			args: args{
				digits: []int{0},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
