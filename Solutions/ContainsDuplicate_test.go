package Solutions

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// Input: nums = [1,2,3,1], k = 3
			// Output: true
			name: "test01",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			// Input: nums = [1,0,1,1], k = 1
			// Output: true
			name: "test02",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			// Input: nums = [1,2,3,1,2,3], k = 2
			// Output: false
			name: "test03",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
