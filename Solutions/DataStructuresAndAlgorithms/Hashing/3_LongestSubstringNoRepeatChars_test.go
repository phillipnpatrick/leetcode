package Solutions

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: s = "abcabcbb"
			// Output: 3
			name: "test01",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			// Input: s = "bbbbb"
			// Output: 1
			name: "test02",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			// Input: s = "pwwkew"
			// Output: 3
			name: "test03",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
