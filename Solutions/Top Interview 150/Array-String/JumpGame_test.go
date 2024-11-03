package Solutions

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{[]int {2,3,1,1,4}},
			want: true,
		},
		{
			name: "test02",
			args: args{[]int {3,2,1,0,4}},
			want: false,
		},
		{
			name: "test03",
			args: args{[]int {0}},
			want: true,
		},
		{
			name: "test04",
			args: args{[]int {0,2,3}},
			want: false,
		},
		{
			name: "test05",
			args: args{[]int {2,0}},
			want: true,
		},
		{
			name: "test06",
			args: args{[]int {1,2,3}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
