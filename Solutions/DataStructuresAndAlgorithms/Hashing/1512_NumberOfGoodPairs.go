package Solutions

// https://leetcode.com/problems/number-of-good-pairs/description/
// Given an array of integers nums, return the number of good pairs.
// A pair (i, j) is called good if nums[i] == nums[j] and i < j.

// Example 1:
// Input: nums = [1,2,3,1,1,3]
// Output: 4
// Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

// Example 2:
// Input: nums = [1,1,1,1]
// Output: 6
// Explanation: Each pair in the array are good.

// Example 3:
// Input: nums = [1,2,3]
// Output: 0

func numIdenticalPairs(nums []int) int {
	frequencies := make(map[int]int)
	pairs := 0

	for i := 0; i < len(nums); i++ {
		frequencies[nums[i]]++
		if frequencies[nums[i]] >= 2 {
			pairs += frequencies[nums[i]] - 1
		}
	}

	return pairs
}
