package Solutions

import (
	"testing"
)

func Test_mycompress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// Example 1:
		// Input: chars = ["a","a","b","b","c","c","c"]
		// Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
		// Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
		{
			name: "test01",
			args: args{chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}},
			want: 6,
			want1: "a2b2c3",
		},
		// Example 2:
		// Input: chars = ["a"]
		// Output: Return 1, and the first character of the input array should be: ["a"]
		// Explanation: The only group is "a", which remains uncompressed since it's a single character.
		{
			name: "test02",
			args: args{chars: []byte{'a'}},
			want: 1,
			want1: "a",
		},
		// Example 3:
		// Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
		// Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
		// Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
		{
			name: "test03",
			args: args{chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}},
			want: 4,
			want1: "ab12",
		},
		// ['a','b','c','d','e','f','1','4','g','1', '8','h','2','i']
		{
			name: "test04",
			args: args{chars: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'h', 'h', 'i'}},
			want: 14,
			want1: "abcdef14g18h2i",
		},
		// ["a","a","a","a","a","b"] => ["a","5","b"]
		{
			name: "test05",
			args: args{chars: []byte{'a', 'a', 'a', 'a', 'a', 'b'}},
			want: 3,
			want1: "a5b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := mycompress(tt.args.chars)
			if got != tt.want {
				t.Errorf("mycompress() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("mycompress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
	}{
		// Example 1:
		// Input: chars = ["a","a","b","b","c","c","c"]
		// Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
		// Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
		{
			name: "test01",
			args: args{chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}},
			want: 6,
		},
		// Example 2:
		// Input: chars = ["a"]
		// Output: Return 1, and the first character of the input array should be: ["a"]
		// Explanation: The only group is "a", which remains uncompressed since it's a single character.
		{
			name: "test02",
			args: args{chars: []byte{'a'}},
			want: 1,
		},
		// Example 3:
		// Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
		// Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
		// Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
		{
			name: "test03",
			args: args{chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}},
			want: 4,
		},
		// ['a','b','c','d','e','f','1','4','g','1', '8','h','2','i']
		{
			name: "test04",
			args: args{chars: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'h', 'h', 'i'}},
			want: 14,
		},
		// ["a","a","a","a","a","b"] => ["a","5","b"]
		{
			name: "test05",
			args: args{chars: []byte{'a', 'a', 'a', 'a', 'a', 'b'}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compress(tt.args.chars)
			if got != tt.want {
				t.Errorf("compress() got = %v, want %v", got, tt.want)
			}
		})
	}
}