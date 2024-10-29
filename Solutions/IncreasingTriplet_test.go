package Solutions

import "testing"

func Test_increasingTriplet(t *testing.T) {
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
			args: args{nums: []int{1,2,3,4,5}},	// Explanation: Any triplet where i < j < k is valid.
			want: true,
		},
		{
			name: "test02",
			args: args{nums: []int{5,4,3,2,1}}, // Explanation: No triplet exists.
			want: false,
		},
		{
			name: "test03",
			args: args{nums: []int{2,1,5,0,4,6}}, // Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
			want: true,
		},
		{
			name: "test04",
			args: args{nums: []int{20,100,10,12,5,13}},
			want: true,
		},
		{
			name: "test05",
			args: args{nums: []int{20,100,10,12,5,7}},
			want: false,
		},
		{
			name: "test06",
			args: args{nums: []int{20,100,10,12,-5,7,-2,-6,13}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %t, want %t", got, tt.want)
			}
		})
	}
}
