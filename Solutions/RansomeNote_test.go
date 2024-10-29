package Solutions

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{ransomNote: "a", magazine: "b"},
			want: false,
		},
		{
			name: "test02",
			args: args{ransomNote: "aa", magazine: "ab"},
			want: false,
		},
		{
			name: "test03",
			args: args{ransomNote: "aa", magazine: "aab"},
			want: true,
		},
		{
			name: "test04",
			args: args{ransomNote: "aa", magazine: "zyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbabcdefghijklmnopqrstuvwxyz"},
			want: true,
		},
		{
			name: "test05",
			args: args{ransomNote: "aab", magazine: "bzyxwvutsrqponmabcdefghijklmnopqrstuvwxyza"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
