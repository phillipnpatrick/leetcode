package Solutions

import "unicode"

// https://leetcode.com/problems/reverse-words-in-a-string-iii/description/
// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Example 1:
//             0123456789
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

// Example 2:
// Input: s = "Mr Ding"
// Output: "rM gniD"

func reverseWords(s string) string {
	runes := make([]rune, len(s))
	left, right := 0, 0

	for right < len(s) {
		if unicode.IsSpace(rune(s[right])) {
			for j := right-1; left < right; j-- {
				runes[left] = rune(s[j])
				left++
			}
			runes[left] = rune(s[right])
			left++
		}
		right++
	}

	for left < len(s) {
		right--
		runes[left] = rune(s[right])
		left++
	}

	return string(runes)
}
