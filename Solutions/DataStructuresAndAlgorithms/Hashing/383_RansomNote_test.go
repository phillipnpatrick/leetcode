package Solutions

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: ransomNote = "a", magazine = "b"
			// Output: false
			name: "test01",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			// Input: ransomNote = "aa", magazine = "ab"
			// Output: false
			name: "test02",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			// Input: ransomNote = "aa", magazine = "aab"
			// Output: true
			name: "test03",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
		{
			name: "test04",
			args: args{
				ransomNote: "aa",
				magazine:   "zyx",
			},
			want: false,
		},
		{
			name: "test05",
			args: args{
				ransomNote: "bg",
				magazine:   "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
