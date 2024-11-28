package Solutions

import "strings"

// https://leetcode.com/problems/permutation-in-string/description/
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.

// Example 1:
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").

// Example 2:
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	if strings.Contains(s2, s1) {
		return true
	}

	f1 := make(map[rune]int)
	for i := 0; i < len(s1); i++ {
		f1[rune(s1[i])]++
	}

	left, right := 0, 0

	for right < len(s2) {
		if len(s1) == right-left+1 {
			f2 := make(map[rune]int)
			for j := left; j <= right; j++ {
				f2[rune(s2[j])]++
			}

			if areMapsEqual(f1, f2) {
				return true
			}

			left++
		}
		right++
	}

	return false
}

func areMapsEqual(m1, m2 map[rune]int) bool {
	for key1, value1 := range m1 {
		if value2, exists := m2[key1]; exists {
			if value1 != value2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
