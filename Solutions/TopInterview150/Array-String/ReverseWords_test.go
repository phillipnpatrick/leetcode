package Solutions

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test01",
			args: args{s: "the sky is blue"},
			want: "blue is sky the",
		},
		{
			name: "test02",
			args: args{s: "  hello world  "},
			want: "world hello",
		},
		{
			name: "test03",
			args: args{s: "a good   example"},
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = '%v'; want => '%v'", got, tt.want)
			}
		})
	}
}
