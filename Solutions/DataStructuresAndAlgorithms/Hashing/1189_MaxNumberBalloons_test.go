package Solutions

import "testing"

func Test_maxNumberofBalloons(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Input: text = "nlaebolko"
			// Output: 1
			name: "test01",
			args: args{
				text: "nlaebolko",
			},
			want: 1,
		},
		{
			// Input: text = "loonbalxballpoon"
			// Output: 2
			name: "test02",
			args: args{
				text: "loonbalxballpoon",
			},
			want: 2,
		},
		{
			// Input: text = "leetcode"
			// Output: 0
			name: "test03",
			args: args{
				text: "leetcode",
			},
			want: 0,
		},
		{
			name: "test04",
			args: args{
				text: "nllbblooon",
			},
			want: 0,
		},
		{
			name: "test05",
			args: args{
				text: "balllllllllllloooooooooon",
			},
			want: 1,
		},
		{
			name: "test06",
			args: args{
				text: "balon",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.args.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
		})
	}
}
