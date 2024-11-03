package Solutions

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: haystack = "sadbutsad", needle = "sad"
			// Output: 0
			name: "test01",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			// Input: haystack = "leetcode", needle = "leeto"
			// Output: -1
			name: "test02",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "test03",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "test04",
			args: args{
				haystack: "aaa",
				needle:   "aaaa",
			},
			want: -1,
		},
		{
			name: "test05",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			name: "test06",
			args: args{
				haystack: "mississippi",
				needle:   "issipi",
			},
			want: -1,
		},
		{
			name: "test07",
			args: args{
				haystack: "mississippi",
				needle:   "sippia",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
