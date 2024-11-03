package Solutions

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test01",
			args: args{matrix: [][]int{{1,2,3},{4,5,6},{7,8,9}}},
			want: [][]int{{7,4,1},{8,5,2},{9,6,3}},
		},
		{
			name: "test02",
			args: args{matrix: [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}},
			want: [][]int{{15,13,2,5},{14,3,4,1},{12,6,8,9},{16,7,10,11}},
		},
		{
			name: "test03",
			args: args{matrix: [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}}},
			want: [][]int{{13,9,5,1},{14,10,6,2},{15,11,7,3},{16,12,8,4}},
		},
		{
			name: "test04",
			args: args{matrix: [][]int{{1,2,3,4,5},{6,7,8,9,10},{11,12,13,14,15},{16,17,18,19,20},{21,22,23,24,25}}},
			want: [][]int{{21,16,11,6,1},{22,17,12,7,2},{23,18,13,8,3},{24,19,14,9,4},{25,20,15,10,5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
