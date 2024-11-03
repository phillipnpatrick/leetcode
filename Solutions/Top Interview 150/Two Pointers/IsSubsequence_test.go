package Solutions

import (
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Input: s = "abc", t = "ahbgdc"
		// Output: true
		{
			name: "test01",
			args: args{s: "abc", t: "ahbgdc"},
			want: true,
		},
		// Input: s = "axc", t = "ahbgdc"
		// Output: false
		{
			name: "test02",
			args: args{s: "axc", t: "ahbgdc"},
			want: false,
		},
		// Input: s = "aec", t = "abcde"
		// Output: false
		{
			name: "test03",
			args: args{s: "aec", t: "abcde"},
			want: false,
		},
		{
			name: "test04",
			args: args{s: "abc", t: "zyxwvuahbgdc"},
			want: true,
		},
		{
			name: "test05",
			args: args{s: "abc", t: "zyxwvutsrqponmlkjihgfed"},
			want: false,
		},
		{
			name: "test06",
			args: args{s: "aaaaaa", t: "bbaaaa"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubsequenceSimple(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Input: s = "abc", t = "ahbgdc"
		// Output: true
		{
			name: "test01",
			args: args{s: "abc", t: "ahbgdc"},
			want: true,
		},
		// Input: s = "axc", t = "ahbgdc"
		// Output: false
		{
			name: "test02",
			args: args{s: "axc", t: "ahbgdc"},
			want: false,
		},
		// Input: s = "aec", t = "abcde"
		// Output: false
		{
			name: "test03",
			args: args{s: "aec", t: "abcde"},
			want: false,
		},
		{
			name: "test04",
			args: args{s: "abc", t: "zyxwvuahbgdc"},
			want: true,
		},
		{
			name: "test05",
			args: args{s: "abc", t: "zyxwvutsrqponmlkjihgfed"},
			want: false,
		},
		{
			name: "test06",
			args: args{s: "aaaaaa", t: "bbaaaa"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequenceSimple(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequenceSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}
