package Solutions

// https://leetcode.com/problems/isomorphic-strings/description/
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
	if len(s) != len(t) {
		return false
	}

	s_t := make(map[byte]byte)
	t_s := make(map[byte]byte)

	for index := 0; index < len(s); index++ {
		stValue, stExists := s_t[s[index]]
		tsValue, tsExists := t_s[t[index]]
		if (stExists && stValue != t[index]) || (tsExists && tsValue != s[index]) {
			return false
		}
		s_t[s[index]] = t[index]
		t_s[t[index]] = s[index]
	}

	return true
}
