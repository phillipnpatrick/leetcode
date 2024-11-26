package Solutions

import "testing"

func Test_closeStrings(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: word1 = "abc", word2 = "bca"
			// Output: true
			name: "test01",
			args: args{
				word1: "abc",
				word2: "bca",
			},
			want: true,
		},
		{
			// Input: word1 = "a", word2 = "aa"
			// Output: false
			name: "test02",
			args: args{
				word1: "a",
				word2: "aa",
			},
			want: false,
		},
		{
			// Input: word1 = "cabbba", word2 = "abbccc"
			// Output: true
			name: "test03",
			args: args{
				word1: "cabbba",
				word2: "abbccc",
			},
			want: true,
		},
		{
			name: "test04",
			args: args{
				word1: "abcde",
				word2: "aecdb",
			},
			want: true,
		},
		{
			name: "test05",
			args: args{
				word1: "aacabb",
				word2: "bbcbaa",
			},
			want: true,
		},
		{
			name: "test06",
			args: args{
				word1: "uau",
				word2: "ssx",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closeStrings(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("closeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
