package Solutions

import (
	"math"
	"strings"
)

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
	max := float64(0)

	// fmt.Printf("len(s) = %d \n", len(s))

	count := -1
	for i := 0; i <= len(s)-k; i++ {
		if count < 0 {
			str := s[i : k+i]
			count = strings.Count(str, "a")
			count += strings.Count(str, "e")
			count += strings.Count(str, "i")
			count += strings.Count(str, "o")
			count += strings.Count(str, "u")
		} else {
			// fmt.Printf("str = %s \t s[%d] = %s \t s[%d] = %s \t", s[i : k+i], i-1, string(s[i-1]), i-1+k, string(s[i-1+k]))
			count -= isVowel(rune(s[i-1]))
			count += isVowel(rune(s[i-1+k]))
			// fmt.Printf("count = %d \n", count)
		}
		max = math.Max(max, float64(count))
	}

	return int(max)
}

func isVowel(letter rune) int {
	if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
		return 1
	}
	return 0
}
