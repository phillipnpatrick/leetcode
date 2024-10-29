package Solutions

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: s = "III"
			// Output: 3
			name: "test01",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			// Input: s = "LVIII"
			// Output: 58
			name: "test02",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			// Input: s = "MCMXCIV"
			// Output: 1994
			name: "test03",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
		{
			name: "test04",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "test05",
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: "test06",
			args: args{
				s: "XL",
			},
			want: 40,
		},
		{
			name: "test07",
			args: args{
				s: "XC",
			},
			want: 90,
		},
		{
			name: "test08",
			args: args{
				s: "CD",
			},
			want: 400,
		},
		{
			name: "test09",
			args: args{
				s: "CM",
			},
			want: 900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
