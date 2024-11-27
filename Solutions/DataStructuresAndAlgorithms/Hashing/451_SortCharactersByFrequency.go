package Solutions

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/sort-characters-by-frequency/description/
// Given a string s, sort it in decreasing order based on the frequency of the characters. The frequency of a character is the number of times it appears in the string.
// Return the sorted string. If there are multiple answers, return any of them.

// Example 1:
// Input: s = "tree"
// Output: "eert"
// Explanation: 'e' appears twice while 'r' and 't' both appear once.
// So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.

// Example 2:
// Input: s = "cccaaa"
// Output: "aaaccc"
// Explanation: Both 'c' and 'a' appear three times, so both "cccaaa" and "aaaccc" are valid answers.
// Note that "cacaca" is incorrect, as the same characters must be together.

// Example 3:
// Input: s = "Aabb"
// Output: "bbAa"
// Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
// Note that 'A' and 'a' are treated as two different characters.

func frequencySort(s string) string {
	frequencies := make(map[rune]int)
	characters := make(map[int]string)
	var indices []int

	for i := 0; i < len(s); i++ {
		frequencies[rune(s[i])]++
	}

	for character, count := range frequencies {
		if _, exists := characters[count]; exists {
			value := characters[count]
			characters[count] = value + buildString(count, character)
		} else {
			characters[count] = buildString(count, character)
		}
	}

	for frequency := range characters {
		indices = append(indices, frequency)
	}

	sort.Ints(indices)

	var builder strings.Builder
	for j := len(indices)-1; j >= 0; j-- {
		builder.WriteString(characters[indices[j]])
	}

	return builder.String()
}

func buildString(repeat int, value rune) string {
	var builder strings.Builder

	for j := 0; j < repeat; j++ {
		builder.WriteString(string(value))
	}

	return builder.String()
}
