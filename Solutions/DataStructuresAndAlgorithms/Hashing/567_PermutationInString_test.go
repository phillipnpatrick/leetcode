package Solutions

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: s1 = "ab", s2 = "eidbaooo"
			// Output: true
			name: "test01",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
		{
			// Input: s1 = "ab", s2 = "eidboaoo"
			// Output: false
			name: "test02",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
