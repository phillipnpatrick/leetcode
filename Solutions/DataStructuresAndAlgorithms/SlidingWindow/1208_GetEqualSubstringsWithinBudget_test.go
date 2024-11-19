package Solutions

import "testing"

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: s = "abcd", t = "bcdf", maxCost = 3
			// Output: 3
			name: "test01",
			args: args{
				s:       "abcd",
				t:       "bcdf",
				maxCost: 3,
			},
			want: 3,
		},
		{
			// Input: s = "abcd", t = "cdef", maxCost = 3
			// Output: 1
			name: "test02",
			args: args{
				s:       "abcd",
				t:       "cdef",
				maxCost: 3,
			},
			want: 1,
		},
		{
			// Input: s = "abcd", t = "acde", maxCost = 0
			// Output: 1
			name: "test03",
			args: args{
				s:       "abcd",
				t:       "acde",
				maxCost: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
