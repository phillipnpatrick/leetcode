package Solutions

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// Input: strs = ["flower","flow","flight"]
			// Output: "fl"
			name: "test01",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			// Input: strs = ["dog","racecar","car"]
			// Output: ""
			name: "test02",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "test03",
			args: args{
				strs: []string{"a"},
			},
			want: "a",
		},
		{
			name: "test04",
			args: args{
				strs: []string{"cir","car"},
			},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
