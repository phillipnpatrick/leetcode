package Solutions

import "math"

// https://leetcode.com/problems/maximum-average-subarray-i/description/

// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10-5 will be accepted.

// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000

func findMaxAverage(nums []int, k int) float64 {
	var avg float64
	left, right, sum := 0, 0, 0

	for right < k {
		sum += nums[right]
		right++
	}

	avg = float64(sum)/float64(k)
	
	for right < len(nums) {
		sum += nums[right] - nums[left]
		avg = math.Max(avg, float64(sum)/float64(k))
		right++
		left++
	}

	return avg
}