package Solutions

// https://leetcode.com/problems/subarray-sum-equals-k/description/
// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,1,1], k = 2
// Output: 2

// Example 2:
// Input: nums = [1,2,3], k = 3
// Output: 2

func subarraySum(nums []int, k int) int {
	counts := make(map[int]int)
	ans, sum := 0, 0

	counts[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans += counts[sum-k]
		counts[sum]++
	}

	return ans
}