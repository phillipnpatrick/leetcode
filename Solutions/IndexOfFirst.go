package Solutions

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Example 1:
// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.

// Example 2:
// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.

func strStr(haystack string, needle string) int {
	index := -1
	if len(needle) <= len(haystack) {
		for i, j := 0, 0; i < len(haystack) && j < len(needle); i++ {
			if needle[j] == haystack[i] {
				if index == -1 {
					index = i
				}
				j++
			} else if index >= 0 {
				index = -1		// reset index to continue search
				i -= j			// backtrack in haystack the length of needle (loop will increase counter + 1)
				j = 0			// reset j to restart search for needle
			}
		}
		if index + len(needle)-1 > len(haystack)-1 {
			index = -1
		}
	}
	return index
}
