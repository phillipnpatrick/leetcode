package Solutions

// https://leetcode.com/problems/minimum-size-subarray-sum/description/
// Given an array of positive integers nums and a positive integer target, return the minimal length of a
// subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Example 1:
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

// Example 2:
// Input: target = 4, nums = [1,4,4]
// Output: 1

// Example 3:
// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

func minSubArrayLen(target int, nums []int) int {
	left, sum, minLength := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		
		for sum >= target {
			sum -= nums[left]
			
			if minLength == 0 {
				minLength = len(nums)
			}
			minLength = min(minLength, right-left+1)
			left++
		}
	}

	return minLength
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
