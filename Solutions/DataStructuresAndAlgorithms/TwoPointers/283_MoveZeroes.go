package Solutions

// https://leetcode.com/problems/move-zeroes/description/
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

// [2,4,6,0,8,0,10,12,0] --> [2,4,6,8,10,12,0,0,0]
func moveZeroes(nums []int) []int {
	i, j := 0, 1

	for j < len(nums) {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		} else if nums[i] != 0 {
			i++
		}
		j++
	}

	return nums
}