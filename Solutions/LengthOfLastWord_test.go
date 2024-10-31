package Solutions

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: s = "Hello World"
			// Output: 5
			name: "test01",
			args: args{
				s: "Hello World",
			},
			want: 5,
		},
		{
			// Input: s = "   fly me   to   the moon  "
			// Output: 4
			name: "test02",
			args: args{
				s: "   fly me   to   the moon  ",
			},
			want: 4,
		},
		{
			// Input: s = "luffy is still joyboy"
			// Output: 6
			name: "test03",
			args: args{
				s: "luffy is still joyboy",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
