package Solutions

import "testing"

func Test_equalPairs(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
			// Output: 1
			name: "test01",
			args: args{
				grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			},
			want: 1,
		},
		{
			// Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
			// Output: 3
			name: "test02",
			args: args{
				grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
			},
			want: 3,
		},
		{
			name: "test03",
			args: args{
				grid: [][]int{{13,13}, {13,13}},
			},
			want: 4,
		},
		{
			name: "test04",
			args: args{
				grid: [][]int{{3,3,3,6,18,3,3,3,3,3},{3,3,3,3,1,3,3,3,3,3},{3,3,3,3,1,3,3,3,3,3},{3,3,3,3,1,3,3,3,3,3},{1,1,1,11,19,1,1,1,1,1},{3,3,3,18,19,3,3,3,3,3},{3,3,3,3,1,3,3,3,3,3},{3,3,3,3,1,3,3,3,3,3},{3,3,3,1,6,3,3,3,3,3},{3,3,3,3,1,3,3,3,3,3}},
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalPairs(tt.args.grid); got != tt.want {
				t.Errorf("equalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
