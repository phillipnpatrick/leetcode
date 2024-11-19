package Solutions

// https://leetcode.com/problems/get-equal-substrings-within-budget/description/
// You are given two strings s and t of the same length and an integer maxCost.
// You want to change s to t. Changing the ith character of s to ith character of t costs |s[i] - t[i]|
// (i.e., the absolute difference between the ASCII values of the characters).
// Return the maximum length of a substring of s that can be changed to be the same as the corresponding substring of t with a cost less than or equal to maxCost.
// If there is no substring from s that can be changed to its corresponding substring from t, return 0.

// Example 1:
// Input: s = "abcd", t = "bcdf", maxCost = 3
// Output: 3
// Explanation: "abc" of s can change to "bcd".
// That costs 3, so the maximum length is 3.

// Example 2:
// Input: s = "abcd", t = "cdef", maxCost = 3
// Output: 1
// Explanation: Each character in s costs 2 to change to character in t,  so the maximum length is 1.

// Example 3:
// Input: s = "abcd", t = "acde", maxCost = 0
// Output: 1
// Explanation: You cannot make any change, so the maximum length is 1.

func equalSubstring(s string, t string, maxCost int) int {
	maximum := 0
	i, j := 0, 0
	cost := 0

	for i < len(s) && j < len(t) {
		cost += absDifference(s[j], t[j])

		for cost > maxCost {
			cost -= absDifference(s[i], t[i])
			i++
		}
		maximum = max(maximum, j-i+1)
		j++
	}

	return maximum
}

func absDifference(a byte, b byte) int {
	n := int(a)
	m := int(b)
	d := n - m

	if d < 0 {
		return 0 - d
	}
	return d
}

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }
