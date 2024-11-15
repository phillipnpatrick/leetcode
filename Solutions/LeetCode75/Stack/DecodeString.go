package Solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Given an encoded string, return its decoded string.
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
// Note that k is guaranteed to be a positive integer.
// You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc.
// Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
// For example, there will not be input like 3a or 2[4].
// The test cases are generated so that the length of the output will never exceed 105.

// Example 1:
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"

// Example 2:
// Input: s = "3[a2[c]]"
// Output: "accaccacc"

// Example 3:
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"

func decodeString(s string) string {
	var sb strings.Builder

	pattern := `(\d+\[[^\]]*\])|([a-zA-Z]+)`
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllString(s, -1)

	// Print the matches
	for _, match := range matches {
		match = formatBrackets(match)

		for strings.ContainsAny(match, "[]") {
			re = regexp.MustCompile(`\d+\[[^\[]*\]`)

			// Find all matches
			innerMatches := re.FindAllString(match, -1)

			// decode matches
			var temp string
			for i := len(innerMatches) - 1; i >= 0; i-- {
				innerMatches[i] = formatBrackets(innerMatches[i])
				temp = fmt.Sprintf("%s%s", temp, decode(innerMatches[i]))
				match = strings.Replace(match, innerMatches[i], temp, 1)
			}
		}
		sb.WriteString(decode(match))
	}

	return sb.String()
}

func formatBrackets(s string) string {
	leftBracket := strings.Count(s, "[")
	rightBracket := strings.Count(s, "]")

	for rightBracket < leftBracket {
		s = fmt.Sprintf("%s]", s)
		rightBracket++
	}

	for rightBracket > leftBracket {
		i := strings.LastIndex(s, "]")
		s = s[:i]
		rightBracket--
	}

	return s
}

func decode(s string) string {
	var sb strings.Builder
	var num int
	var encoded string

	for len(s) > 0 {
		leftBracket := strings.Index(s, "[")
		rightBracket := strings.Index(s, "]")

		if leftBracket < 0 && rightBracket < 0 {
			num = 1
			encoded = s
			s = ""
		} else {
			num, _ = strconv.Atoi(s[:leftBracket])
			encoded = s[leftBracket+1 : rightBracket]
			s = s[rightBracket+1:]
		}

		for i := 0; i < num; i++ {
			sb.WriteString(encoded)
		}
	}

	return sb.String()
}
