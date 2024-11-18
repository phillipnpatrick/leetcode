package Solutions

import "sort"

// https://leetcode.com/problems/string-compression/?envType=study-plan-v2&envId=leetcode-75

// You are given an integer array nums and an integer k.
// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
// Return the maximum number of operations you can perform on the array.

// Example 1:
// Input: nums = [1,2,3,4], k = 5
// Output: 2
// Explanation: Starting with nums = [1,2,3,4]:
// - Remove numbers 1 and 4, then nums = [2,3]
// - Remove numbers 2 and 3, then nums = []
// There are no more pairs that sum up to 5, hence a total of 2 operations.

// Example 2:
// Input: nums = [3,1,3,4,3], k = 6
// Output: 1
// Explanation: Starting with nums = [3,1,3,4,3]:
// - Remove the first two 3's, then nums = [1,4,3]
// There are no more pairs that sum up to 6, hence a total of 1 operation.

// nums: []int{4,4,1,3,1,3,2,2,5,5,1,5,2,1,2,3,5,4},
// k:    2,
// nums: []int{2,2,2,3,1,1,4,1},
// k:    4,
func maxOperations(nums []int, k int) int {
	operations := 0
	
	// for left := 0; left < len(nums); left++ {
	// 	if nums[left] < k {
	// 		for right := left+1; right < len(nums); right++ {
	// 			if nums[left] + nums[right] == k {
	// 				operations++
	// 				nums = append(nums[:right], nums[right+1:]...)
	// 				nums = append(nums[:left], nums[left+1:]...)
	// 				left--	// adjust index for removal
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	operations = cheating(nums, k)

	return operations
}

func cheating(nums[]int, k int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	operation := 0

	for left < right {
		if nums[left] + nums[right] == k {
			operation++
			left++
			right--
		} else if nums[left] + nums[right] < k {
			left++
		} else {
			right--
		}
	}

	return operation
}