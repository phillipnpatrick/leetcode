package Solutions

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: s = "()"
			// Output: true
			name: "test01",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			// Input: s = "()[]{}"
			// Output: true
			name: "test02",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			// Input: s = "(]"
			// Output: false
			name: "test03",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			// Input: s = "([])"
			// Output: true
			name: "test04",
			args: args{
				s: "([])",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
