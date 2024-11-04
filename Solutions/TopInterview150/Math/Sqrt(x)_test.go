package Solutions

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: x = 4
			// Output: 2
			name: "test01",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			// Input: x = 8
			// Output: 2
			name: "test02",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "test03",
			args: args{
				x: 9,
			},
			want: 3,
		},
		{
			name: "test04",
			args: args{
				x: 13,
			},
			want: 3,
		},
		{
			name: "test05",
			args: args{
				x: 16,
			},
			want: 4,
		},
		{
			name: "test06",
			args: args{
				x: 90,
			},
			want: 9,
		},
		{
			name: "test07",
			args: args{
				x: 100,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
