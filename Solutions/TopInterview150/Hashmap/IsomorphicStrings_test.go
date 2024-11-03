package Solutions

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: s = "egg", t = "add"
			// Output: true
			name: "test01",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			// Input: s = "foo", t = "bar"
			// Output: false
			name: "test02",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			// Input: s = "paper", t = "title"
			// Output: true
			name: "test03",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
		{
			name: "test04",
			args: args{
				s: "badc",
				t: "baba",
			},
			want: false,
		},
		{
			name: "test05",
			args: args{
				s: "abab",
				t: "baba",
			},
			want: true,
		},
		{
			name: "test06",
			args: args{
				s: "bbbaaaba",
				t: "aaabbbba",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
