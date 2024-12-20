package Solutions

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: jewels = "aA", stones = "aAAbbbb"
			// Output: 3
			name: "test01",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			// Input: jewels = "z", stones = "ZZ"
			// Output: 0
			name: "test02",
			args: args{
				jewels: "z",
				stones: "ZZ",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
