package Solutions

import "testing"

func Test_countElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: arr = [1,2,3]
			// Output: 2
			name: "test01",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			// Input: arr = [1,1,3,3,5,5,7,7]
			// Output: 0
			name: "test02",
			args: args{
				arr: []int{1, 1, 3, 3, 5, 5, 7, 7},
			},
			want: 0,
		},
		{
			name: "test03",
			args: args{
				arr: []int{1,3,2,3,5,0},
			},
			want: 3,
		},
		{
			name: "test04",
			args: args{
				arr: []int{1,1,2,2},
			},
			want: 2,
		},
		{
			name: "test05",
			args: args{
				arr: []int{1,1,2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.arr); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
