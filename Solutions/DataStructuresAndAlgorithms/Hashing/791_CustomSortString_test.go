package Solutions

import "testing"

func Test_customSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// Input: order = "cba", s = "abcd"
			// Output: "cbad"
			name: "test01",
			args: args{
				order: "cba",
				s:     "abcd",
			},
			want: "cbad",
		},
		{
			// Input: order = "bcafg", s = "abcd"
			// Output: "bcad"
			name: "test02",
			args: args{
				order: "bcafg",
				s:     "abcd",
			},
			want: "bcad",
		},
		{
			name: "test03",
			args: args{
				order: "ab",
				s:     "acbbadbba",
			},
			want: "aaabbbbcd",
		},
		{
			name: "test04",
			args: args{
				order: "kqep",
				s:     "pekeq",
			},
			want: "kqeep",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := customSortString(tt.args.order, tt.args.s); got != tt.want {
				t.Errorf("customSortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
