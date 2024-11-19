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
			// Input: s = "Let's take LeetCode contest"
			// Output: "s'teL ekat edoCteeL tsetnoc"
			name: "test01",
			args: args{
				s: "Let's take LeetCode contest",
			},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			// Input: s = "Mr Ding"
			// Output: "rM gniD"
			name: "test02",
			args: args{
				s: "Mr Ding",
			},
			want: "rM gniD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
