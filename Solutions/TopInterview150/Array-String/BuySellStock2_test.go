package Solutions

import "testing"

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{[]int {7,1,5,3,6,4}},
			want: 7,
		},
		{
			name: "test02",
			args: args{[]int {1,2,3,4,5}},
			want: 4,
		},
		{
			name: "test03",
			args: args{[]int {7,6,4,3,1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
