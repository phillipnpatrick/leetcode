package Solutions

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/count-number-of-nice-subarrays/description/
// Given an array of integers nums and an integer k. A continuous subarray is called nice if there are k odd numbers on it.
// Return the number of nice sub-arrays.

// Example 1:
// Input: nums = [1,1,2,1,1], k = 3
// Output: 2
// map: [0: 1, 1: 1, 2: 2, 3: 1, 4: 1]
// Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].

// Example 2:
// Input: nums = [2,4,6], k = 1
// Output: 0
// Explanation: There are no odd numbers in the array.

// Example 3:
// Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
// map: [0: 3, 1: 3, 2: 4]
// Output: 16

func numberOfSubarrays(nums []int, k int) int {
	ans, sum := 0, 0
	counts := make(map[int]int)	// key = sum, value is count of that sum
	var sums, anss strings.Builder

	counts[0] = 1	// this assures that the sum - k gets counted
	for i := 0; i < len(nums); i++ {
		sum += nums[i]%2		// only adds 1 when array value is odd
		ans += counts[sum-k]	// if counts[sum-k] doesn't exists, add 0;
		counts[sum]++			// adds new sum to map
		sums.WriteString(fmt.Sprintf("sum: %d, ", sum))
		anss.WriteString(fmt.Sprintf("ans: %d, ", ans))
	}
	fmt.Printf("%s\n", sums.String())
	fmt.Printf("%s\n", anss.String())

	return ans
}
