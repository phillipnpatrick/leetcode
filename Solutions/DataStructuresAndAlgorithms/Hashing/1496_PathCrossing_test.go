package Solutions

import "testing"

func Test_isPathCrossing(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: path = "NES"
			// Output: false
			name: "test01",
			args: args{
				path: "NES",
			},
			want: false,
		},
		{
			// Input: path = "NESWW"
			// Output: true
			name: "test02",
			args: args{
				path: "NESWW",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPathCrossing(tt.args.path); got != tt.want {
				t.Errorf("isPathCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
