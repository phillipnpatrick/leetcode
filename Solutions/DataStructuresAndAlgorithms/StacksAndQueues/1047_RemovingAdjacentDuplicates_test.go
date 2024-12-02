package Solutions

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// Input: s = "abbaca"
			// Output: "ca"
			name: "test01",
			args: args{
				s: "abbaca",
			},
			want: "ca",
		},
		{
			// Input: s = "azxxzy"
			// Output: "ay"
			name: "test02",
			args: args{
				s: "azxxzy",
			},
			want: "ay",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
