package Solutions

import (
	"unicode"
)

// https://leetcode.com/problems/reverse-only-letters/description/
// Given a string s, reverse the string according to the following rules:

// All the characters that are not English letters remain in the same position.
// All the English letters (lowercase or uppercase) should be reversed.
// Return s after reversing it.

// Example 1:
// Input: s = "ab-cd"
// Output: "dc-ba"

// Example 2:
// Input: s = "a-bC-dEf-ghIj"
// Output: "j-Ih-gfE-dCba"

// Example 3:
// Input: s = "Test1ng-Leet=code-Q!"
// Output: "Qedo1ct-eeLg=ntse-T!"

func reverseOnlyLetters(s string) string {
	characters := []rune(s)
	left, right := 0, len(s)-1

	for left < right {
		if !unicode.IsLetter(rune(s[left])) {
			left++
		}
		if !unicode.IsLetter(rune(s[right])) {
			right--
		}

		if unicode.IsLetter(rune(s[left])) && unicode.IsLetter(rune(s[right])) {
			characters[left], characters[right] = characters[right], characters[left]
			left++
			right--
		}		
	}

	return string(characters)
}
