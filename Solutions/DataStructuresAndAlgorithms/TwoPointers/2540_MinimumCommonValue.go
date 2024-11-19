package Solutions

import "fmt"

// https://leetcode.com/problems/minimum-common-value/description/
// Given two integer arrays nums1 and nums2, sorted in non-decreasing order, return the minimum integer common to both arrays.
// If there is no common integer amongst nums1 and nums2, return -1.
// Note that an integer is said to be common to nums1 and nums2 if both arrays have at least one occurrence of that integer.

// Example 1:
// Input: nums1 = [1,2,3], nums2 = [2,4]
// Output: 2
// Explanation: The smallest element common to both arrays is 2, so we return 2.

// Example 2:
// Input: nums1 = [1,2,3,6], nums2 = [2,3,4,5]
// Output: 2
// Explanation: There are two common elements in the array 2 and 3 out of which 2 is the smallest, so 2 is returned.

func getCommon(nums1 []int, nums2 []int) int {
	common := -1
	right1 := len(nums1) - 1
	right2 := len(nums2) - 1

	fmt.Printf("nums1: %v\n", nums1)
	fmt.Printf("nums2: %v\n", nums2)

	for right1 >= 0 && right2 >= 0 {
		fmt.Printf("nums1[%d] = %d \t nums2[%d] = %d \t ", right1, nums1[right1], right2, nums2[right2])
		if nums1[right1] > nums2[right2] {
			right1--
		} else if nums2[right2] > nums1[right1] {
			right2--
		} else if nums1[right1] == nums2[right2] {
			common = nums1[right1]
			right1--
			right2--
		}
		fmt.Printf("common = %d \n", common)
	}

	return common
}
