package Solutions

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{[]int {2,3,1,1,4}},
			want: 2,
		},
		{
			name: "test02",
			args: args{[]int {2,3,0,1,4}},
			want: 2,
		},
		{
			name: "test03",
			args: args{[]int {0}},
			want: 0,
		},
		{
			name: "test04",
			args: args{[]int {1,2,1,1,1}},
			want: 3,
		},
		{
			name: "test05",
			args: args{[]int {2,1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
