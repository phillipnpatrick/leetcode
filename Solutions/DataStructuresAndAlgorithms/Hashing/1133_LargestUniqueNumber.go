package Solutions

// https://leetcode.com/problems/largest-unique-number/description/
// Given an integer array nums, return the largest integer that only occurs once. If no integer occurs once, return -1.

// Example 1:
// Input: nums = [5,7,3,9,4,9,8,3,1]
// Output: 8
// Explanation: The maximum integer in the array is 9 but it is repeated. The number 8 occurs only once, so it is the answer.

// Example 2:
// Input: nums = [9,9,8,8]
// Output: -1
// Explanation: There is no number that occurs only once.

func largestUniqueNumber(nums []int) int {
    once := -1
	counts := make(map[int]int)

	for i :=0; i < len(nums); i++ {
		counts[nums[i]]++
	}

	for key, value := range counts {
		if value == 1 {
			if key > once {
				once = key
			}
		}
	}

	return once
}