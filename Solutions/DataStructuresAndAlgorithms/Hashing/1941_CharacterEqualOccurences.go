package Solutions

// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/description/
// Given a string s, return true if s is a good string, or false otherwise.
// A string s is good if all the characters that appear in s have the same number of occurrences (i.e., the same frequency).

// Example 1:
// Input: s = "abacbc"
// Output: true
// Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.

// Example 2:
// Input: s = "aaabb"
// Output: false
// Explanation: The characters that appear in s are 'a' and 'b'.
// 'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.

func areOccurrencesEqual(s string) bool {
	counts := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	// Add the frequency of each character to a set
	// The set should only have one value if all the characters have the same frequency
	frequencies := make(map[int]struct{})
	for _, value := range counts {
		frequencies[value] = struct{}{}
	}

	return len(frequencies) == 1
}
