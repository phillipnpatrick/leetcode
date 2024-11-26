package Solutions

import (
	"sort"
)

// https://leetcode.com/problems/determine-if-two-strings-are-close/?envType=study-plan-v2&envId=leetcode-75

// Two strings are considered close if you can attain one from the other using the following operations:
// Operation 1: Swap any two existing characters.
// For example, abcde -> aecdb
// Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
// For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
// You can use the operations on either string as many times as necessary.
// Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

// Example 1:
// Input: word1 = "abc", word2 = "bca"
// Output: true
// Explanation: You can attain word2 from word1 in 2 operations.
// Apply Operation 1: "abc" -> "acb"
// Apply Operation 1: "acb" -> "bca"

// Example 2:
// Input: word1 = "a", word2 = "aa"
// Output: false
// Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.

// Example 3:
// Input: word1 = "cabbba", word2 = "abbccc"
// Output: true
// Explanation: You can attain word2 from word1 in 3 operations.
// Apply Operation 1: "cabbba" -> "caabbb"
// Apply Operation 2: "caabbb" -> "baaccc"
// Apply Operation 2: "baaccc" -> "abbccc"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := getCharFrequency(word1)
	freq2 := getCharFrequency(word2)

	if !haveSameKeys(freq1, freq2) {
		return false
	}

	f1 := make([]int, len(word1))
	f2 := make([]int, len(word2))

	index := 0
	for _, freq := range freq2 {
		f2[index] = freq
		index++
	}

	index = 0
	for _, freq := range freq1 {
		f1[index] = freq
		index++
	}

	return haveSameValues(f1, f2)
}

func getCharFrequency(word string) map[rune]int {
	m := make(map[rune]int)

	// get frequency of each character
	for _, char := range word {
		m[char]++
	}

	return m
}

func haveSameKeys(map1, map2 map[rune]int) bool {
	// Check if both maps have the same number of keys
	if len(map1) != len(map2) {
		return false
	}

	// Check if all keys in map1 are in map2
	for key := range map1 {
		if _, exists := map2[key]; !exists {
			return false
		}
	}

	// Since the lengths are the same, if all keys in map1 exist in map2,
	// then the maps have the same keys
	return true
}

func haveSameValues(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	sort.Ints(s1)
	sort.Ints(s2)

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}

	}
	return true
}
