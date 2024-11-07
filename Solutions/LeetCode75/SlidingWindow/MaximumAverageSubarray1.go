package Solutions

import "math"

// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10^(-5) will be accepted.

// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000

func findMaxAverage(nums []int, k int) float64 {
	maxAverage := -math.MaxFloat64
	sum := float64(0)

	for i := 0; i <= len(nums)-k; i++ {
		if i == 0 {
			for j := i; j-i < k; j++ {
				sum += float64(nums[j])
			}
		} else {
			sum -= float64(nums[i-1])
			sum += float64(nums[i+k-1])
		}
		average := float64(sum / float64(k))
		maxAverage = math.Max(average, maxAverage)
	}

	return maxAverage
}
