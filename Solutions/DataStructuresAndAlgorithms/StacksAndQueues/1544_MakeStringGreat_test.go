package Solutions

import "testing"

func Test_makeGood(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// Input: s = "leEeetcode"
			// Output: "leetcode"
			name: "test01",
			args: args{
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			// Input: s = "abBAcC"
			// Output: ""
			name: "test02",
			args: args{
				s: "abBAcC",
			},
			want: "",
		},
		{
			// Input: s = "s"
			// Output: "s"
			name: "test03",
			args: args{
				s: "s",
			},
			want: "s",
		},
		{
			name: "test04",
			args: args{
				s: "Pp",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.args.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
