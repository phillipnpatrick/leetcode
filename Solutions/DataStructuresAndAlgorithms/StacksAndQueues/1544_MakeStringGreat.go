package Solutions

import ("unicode"
"strings")

// https://leetcode.com/problems/make-the-string-great/description/
// Given a string s of lower and upper case English letters.

// A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:
// 	0 <= i <= s.length - 2
// 	s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.
// To make the string good, you can choose two adjacent characters that make the string bad and remove them. You can keep doing this until the string becomes good.
// Return the string after making it good. The answer is guaranteed to be unique under the given constraints.
// Notice that an empty string is also good.

// Example 1:
// Input: s = "leEeetcode"
// Output: "leetcode"
// Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".

// Example 2:
// Input: s = "abBAcC"
// Output: ""
// Explanation: We have many possible scenarios, and all lead to the same answer. For example:
// "abBAcC" --> "aAcC" --> "cC" --> ""
// "abBAcC" --> "abBA" --> "aA" --> ""

// Example 3:
// Input: s = "s"
// Output: "s"

func makeGood(s string) string {
	input := []rune(s)
	stack := make([]rune, len(input)+1)

	for i := 0; i < len(input); i++ {
		c := input[i]
		peek := stack[len(stack)-1]

		if isLowerUpper(c, peek) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return toString(stack)
}

func isLowerUpper(c rune, peek rune) bool {
	if unicode.IsUpper(peek) && c == unicode.ToLower(peek) {
		return true
	}

	if unicode.IsLower(peek) && c == unicode.ToUpper(peek) {
		return true
	}

	return false
}

func toString(stack []rune) string {
	var sb strings.Builder
	for i := 0; i < len(stack); i++ {
		if stack[i] != 0 {
			sb.WriteRune(stack[i])
		}
	}

	return sb.String()
}