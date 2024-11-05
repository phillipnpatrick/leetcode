package Solutions

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: x = 121
			// Output: true
			name: "test01",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			// Input: x = -121
			// Output: false
			name: "test02",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			// Input: x = 10
			// Output: false
			name: "test03",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "test04",
			args: args{
				x: 123321,
			},
			want: true,
		},
		{
			name: "test05",
			args: args{
				x: 1234321,
			},
			want: true,
		},
		{
			name: "test06",
			args: args{
				x: 1234567,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
