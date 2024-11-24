package Solutions

// https://leetcode.com/problems/contiguous-array/description/
// Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.

// Example 1:
// Input: nums = [0,1] = (2+4)/3 = 6/3 = 2
// Output: 2
// Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.

// Example 2:
// Input: nums = [0,1,0] = (2+4+2)/3 = 8/3 = 2
// Output: 2
// Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

// [0,0,1,1] = (2+2+4+4)/3 = 12/3 = 4
func findMaxLength(nums []int) int {
	// Map to store counts and their first occurrence indices
	countMap := make(map[int]int)
	countMap[0] = -1 // Initialize with a base case for count = 0

	maxLen := 0
	count := 0

	for i, num := range nums {
		// Update count based on the current number
		if num == 1 {
			count++
		} else {
			count--
		}

		// Check if this count has been seen before
		if index, exists := countMap[count]; exists {
			// Update maxLen with the length of the subarray
			maxLen = max(maxLen, i-index)
		} else {
			// Store the first occurrence of this count
			countMap[count] = i
		}
	}

	return maxLen
}

// public class Solution {

//     public int findMaxLength(int[] nums) {
//         Map<Integer, Integer> map = new HashMap<>();
//         map.put(0, -1);
//         int maxlen = 0, count = 0;
//         for (int i = 0; i < nums.length; i++) {
//             count = count + (nums[i] == 1 ? 1 : -1);
//             if (map.containsKey(count)) {
//                 maxlen = Math.max(maxlen, i - map.get(count));
//             } else {
//                 map.put(count, i);
//             }
//         }
//         return maxlen;
//     }
// }
