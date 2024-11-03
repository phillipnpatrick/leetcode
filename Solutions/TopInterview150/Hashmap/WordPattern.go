package Solutions

import (
	"strings"
)

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
	letters := []rune(pattern)
	words := strings.Fields(s)

	if len(letters) != len(words) {
		return false
	}
	
	lettersToWords := make(map[rune]string)
	wordsToLetters := make(map[string]rune)

	for i := 0; i < len(letters); i++ {
		lettersToWords[letters[i]] = words[i]
		wordsToLetters[words[i]] = letters[i]
	}

	var patternToS, sToPattern strings.Builder
	for j := 0; j < len(letters); j++ {
		patternToS.WriteString(lettersToWords[letters[j]])
		sToPattern.WriteRune(wordsToLetters[words[j]])
	}

	sWithoutSpaces := strings.Join(words, "")

	return patternToS.String() == sWithoutSpaces && sToPattern.String() == pattern
}