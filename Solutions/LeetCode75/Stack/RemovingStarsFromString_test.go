package Solutions

import "testing"

func Test_removeStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// Input: s = "leet**cod*e"
			// Output: "lecoe"
			name: "test01",
			args: args{
				s: "leet**cod*e",
			},
			want: "lecoe",
		},
		{
			// Input: s = "erase*****"
			// Output: ""
			name: "test02",
			args: args{
				s: "erase*****",
			},
			want: "",
		},
		{
			name: "test03",
			args: args{
				s: "abb*cdfg*****x*",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeStars(tt.args.s); got != tt.want {
				t.Errorf("removeStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
