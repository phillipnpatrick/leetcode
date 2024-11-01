package Solutions

import "strings"

// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters.
// No two characters may map to the same character, but a character may map to itself.

// Example 1:
// Input: s = "egg", t = "add"
// Output: true
// Explanation:
// The strings s and t can be made identical by:
// Mapping 'e' to 'a'.
// Mapping 'g' to 'd'.

// Example 2:
// Input: s = "foo", t = "bar"
// Output: false
// Explanation:
// The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.

// Example 3:
// Input: s = "paper", t = "title"
// Output: true

func isIsomorphic(s string, t string) bool {
	mapTtoS := make(map[string]string)
	mapStoT := make(map[string]string)

	for i, j := 0, 0; i < len(s) && j < len(t); i, j = i+1, j+1 {
		mapTtoS[string(t[i])] = string(s[j])
		mapStoT[string(s[j])] = string(t[i])
	}
	
	newS := replace(s, mapStoT)
	newT := replace(t, mapTtoS)

	return s == newT && t == newS
}

func replace(s string, isoMap map[string]string) string {
	var output strings.Builder
	for k := 0; k < len(s); k++ {
		value, exists := isoMap[string(s[k])]
		if exists {
			output.WriteString(value)
		}
	}

	return output.String()
}