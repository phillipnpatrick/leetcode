package Solutions

import "testing"

func Test_checkIfPangram(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
			// Output: true
			name: "test01",
			args: args{
				sentence: "thequickbrownfoxjumpsoverthelazydog",
			},
			want: true,
		},
		{
			// Input: sentence = "leetcode"
			// Output: false
			name: "test02",
			args: args{
				sentence: "leetcode",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPangram(tt.args.sentence); got != tt.want {
				t.Errorf("checkIfPangram() = %v, want %v", got, tt.want)
			}
		})
	}
}
