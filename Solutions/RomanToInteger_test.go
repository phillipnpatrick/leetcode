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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
