package Solutions

import "strings"

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
// Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

// Example 1:
// Input: s = "abciiidef", k = 3
// Output: 3
// Explanation: The substring "iii" contains 3 vowel letters.

// Example 2:
// Input: s = "aeiou", k = 2
// Output: 2
// Explanation: Any substring of length 2 contains 2 vowels.

// Example 3:
// Input: s = "leetcode", k = 3
// Output: 2
// Explanation: "lee", "eet" and "ode" contain 2 vowels.

func maxVowels(s string, k int) int {
	count, left, maximum := 0, 0, 0

	for right := 0; right < len(s); right++ {
		if isVowel(s[right]) {
			count++
		}
		for right-left >= k {
			if isVowel(s[left]) {
				count--
			}
			left++
		}

		maximum = max(maximum, count)
	}

	return maximum
}

func isVowel(value byte) bool {
	return strings.ContainsAny(string(value), "aeiou")
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
