package Solutions

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			// Input: s = "tree"
			// Output: "eert"
			name: "test01",
			args: args{
				s: "tree",
			},
			want: []string{"eert", "eetr"},
		},
		{
			// Input: s = "cccaaa"
			// Output: "aaaccc"
			name: "test02",
			args: args{
				s: "cccaaa",
			},
			want: []string{"aaaccc", "cccaaa"},
		},
		{
			// Input: s = "Aabb"
			// Output: "bbAa"
			name: "test03",
			args: args{
				s: "Aabb",
			},
			want: []string{"bbAa", "bbaA"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := frequencySort(tt.args.s)

			if !stringInSlice(got, tt.want) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringInSlice(target string, slice []string) bool {
	return func() bool {
		for _, str := range slice {
			if str == target {
				return true
			}
		}
		return false
	}()
}
