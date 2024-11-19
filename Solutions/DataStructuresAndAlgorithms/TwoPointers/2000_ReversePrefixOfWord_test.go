package Solutions

import "testing"

func Test_reversePrefix(t *testing.T) {
	type args struct {
		word string
		ch   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// Input: word = "abcdefd", ch = "d"
			// Output: "dcbaefd"
			name: "test01",
			args: args{
				word: "abcdefd",
				ch:   'd',
			},
			want: "dcbaefd",
		},
		{
			// Input: word = "xyxzxe", ch = "z"
			// Output: "zxyxxe"
			name: "test02",
			args: args{
				word: "xyxzxe",
				ch:   'z',
			},
			want: "zxyxxe",
		},
		{
			// Input: word = "abcd", ch = "z"
			// Output: "abcd"
			name: "test03",
			args: args{
				word: "abcd",
				ch:   'z',
			},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefix(tt.args.word, tt.args.ch); got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
