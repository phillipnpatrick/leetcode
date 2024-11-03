package Solutions

import "testing"

func TestReverseVowels(t *testing.T) {
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
			args: args{s: "IceCreAm"},
			want: "AceCreIm",
		},
		{
			name: "test02",
			args: args{s: "leetcode"},
			want: "leotcede",
		},
		{
			name: "test03",
			args: args{s: " "},
			want: " ",
		},
		{
			name: "test04",
			args: args{s: "a."},
			want: "a.",
		},
		{
			name: "test05",
			args: args{s: "!!!"},
			want: "!!!",
		},
		{
			name: "test06",
			args: args{s: "dad"},
			want: "dad",
		},
		{
			name: "test07",
			args: args{s: "dadedid"},
			want: "didedad",
		},
		{
			name: "test08",
			args: args{s: "Euston saw I was not Sue."},
			want: "euston saw I was not SuE.",
		},
		{
			name: "test09",
			args: args{s: "Marge, let's \"went.\" I await news telegram."},
			want: "Marge, let's \"went.\" i awaIt news telegram.",
		},
		{
			name: "test10",
			args: args{s: "Live on evasions? No, I save no evil."},
			want: "Live on evasIons? No, i save no evil.",
		},
		{
			name: "test11",
			args: args{s: "Resume so pacific a pose, muser."},
			want: "Resume so pacific a pose, muser.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseVowels(tt.args.s); got != tt.want {
				t.Errorf("%v: ReverseVowels() = \"%v\", want \"%v\"", tt.name, got, tt.want)
			}
		})
	}
}
