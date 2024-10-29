package Solutions

import "testing"

func Test_maxProfit(t *testing.T) {
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
			want: 5,
		},
		{
			name: "test02",
			args: args{[]int {7,6,4,3,1}},
			want: 0,
		},
		{
			name: "test03",
			args: args{[]int {2,4,1}},
			want: 2,
		},
		{
			name: "test04",
			args: args{[]int {200,400,10,300,5,40}},
			want: 290,
		},
		{
			name: "test05",
			args: args{[]int {200,400,10,300,500,40000}},
			want: 39990,
		},
		{
			name: "test06",
			args: args{[]int {200,400,10,300,500,40000,60000,80000,100000}},
			want: 99990,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
