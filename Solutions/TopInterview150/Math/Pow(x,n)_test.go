package Solutions

import (
	"testing"
)

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			// Input: x = 2.00000, n = 10
			// Output: 1024.00000
			name: "test01",
			args: args{
				x: 2.00000,
				n: 10,
			},
			want: 1024.00000,
		},
		{
			// Input: x = 2.10000, n = 3
			// Output: 9.26100
			name: "test02",
			args: args{
				x: 2.10000,
				n: 3,
			},
			want: 9.26100,
		},
		{
			// Input: x = 2.00000, n = -2
			// Output: 0.25000
			name: "test03",
			args: args{
				x: 2.00000,
				n: -2,
			},
			want: 0.25000,
		},
		{
			name: "test04",
			args: args{
				x: -2.00000,
				n: 3,
			},
			want: -8.00000,
		},
		{
			name: "test05",
			args: args{
				x: -2.00000,
				n: 4,
			},
			want: 16.00000,
		},
		{
			name: "test06",
			args: args{
				x: 8.95371,
				n: -1,
			},
			want: 0.11169,
		},
		{
			name: "test07",
			args: args{
				x: 34.00515,
				n: -3,
			},
			want: 3e-05,
		},
		{
			name: "test08",
			args: args{
				x: 0.00001,
				n: 2147483647,
			},
			want: 0,
		},
		{
			name: "test09",
			args: args{
				x: 3.76050,
				n: -8,
			},
			want: 3e-05,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
