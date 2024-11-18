package Solutions

import "testing"

func Test_uniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: arr = [1,2,2,1,1,3]
			// Output: true
			name: "test01",
			args: args{
				arr: []int{1, 2, 2, 1, 1, 3},
			},
			want: true,
		},
		{
			// Input: arr = [1,2]
			// Output: false
			name: "test02",
			args: args{
				arr: []int{1, 2},
			},
			want: false,
		},
		{
			// Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
			// Output: true
			name: "test03",
			args: args{
				arr: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
