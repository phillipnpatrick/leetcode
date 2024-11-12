package Solutions

// Given a binary array nums, you should delete one element from it.
// Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

// Example 1:
// Input: nums = [1,1,0,1]
// Output: 3
// Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.

// Example 2:
// Input: nums = [0,1,1,1,0,1,1,0,1]
// Output: 5
// Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].

// Example 3:
// Input: nums = [1,1,1]
// Output: 2
// Explanation: You must delete one element.

func longestSubarray(nums []int) int {
	// see solution for MaxConsecutiveOnes
	k := 1	// number of zeros allowed in subarray
    left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			k -= 1
		}

		if k < 0 {
			if nums[left] == 0 {
				k += 1
			}
			left += 1
		}
		right++
	}

	return right - left - 1
}