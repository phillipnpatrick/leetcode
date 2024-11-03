package Solutions

import "math"

// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

// Example 1:
// Input: nums = [1,2,3,1], k = 3
// Output: true

// Example 2:
// Input: nums = [1,0,1,1], k = 1
// Output: true

// Example 3:
// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false

func containsNearbyDuplicate(nums []int, k int) bool {
	indices := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if len(indices) > 0 {
			j := indices[nums[i]]

			if nums[i] == nums[j] && int(math.Abs(float64(i-j))) <= k {
				return true
			}
		}
		indices[nums[i]] = i
	}

	return false
}
