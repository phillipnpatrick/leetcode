package Solutions

import "strings"

// https://leetcode.com/problems/word-pattern/description/
// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:
// Each letter in pattern maps to exactly one unique word in s.
// Each unique word in s maps to exactly one letter in pattern.
// No two letters map to the same word, and no two words map to the same letter.

// Example 1:
// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true
// Explanation: The bijection can be established as:
// 'a' maps to "dog".
// 'b' maps to "cat".

// Example 2:
// Input: pattern = "abba", s = "dog cat cat fish"
// Output: false

// Example 3:
// Input: pattern = "aaaa", s = "dog cat cat dog"
// Output: false

func wordPattern(pattern string, s string) bool {
	values := strings.Split(s, " ")
	if len(values) != len(pattern) {
		return false
	}

	pattern_s := make(map[byte]string)
	s_pattern := make(map[string]byte)

	for index := 0; index < len(pattern); index++ {
		patternSValue, patternSExists := pattern_s[pattern[index]]
		sPatternValue, sPatternExists := s_pattern[values[index]]
		if (patternSExists && patternSValue != values[index]) || (sPatternExists && sPatternValue != pattern[index]) {
			return false
		}
		pattern_s[pattern[index]] = values[index]
		s_pattern[values[index]] = pattern[index]
	}

	return true
}