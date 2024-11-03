package Solutions

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test01",
			args: args{
				num: 3749,
			},
			want: "MMMDCCXLIX",
		},
		{
			name: "test02",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "test03",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "test04",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "test05",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "test06",
			args: args{
				num: 40,
			},
			want: "XL",
		},
		{
			name: "test07",
			args: args{
				num: 90,
			},
			want: "XC",
		},
		{
			name: "test08",
			args: args{
				num: 400,
			},
			want: "CD",
		},
		{
			name: "test09",
			args: args{
				num: 900,
			},
			want: "CM",
		},
		{
			name: "test10",
			args: args{
				num: 999,
			},
			want: "CMXCIX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
