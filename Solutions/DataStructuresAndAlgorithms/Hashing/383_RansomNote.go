package Solutions

// https://leetcode.com/problems/ransom-note/description
// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.

// Example 1:
// Input: ransomNote = "a", magazine = "b"
// Output: false

// Example 2:
// Input: ransomNote = "aa", magazine = "ab"
// Output: false

// Example 3:
// Input: ransomNote = "aa", magazine = "aab"
// Output: true

func canConstruct(ransomNote string, magazine string) bool {
	note, mag := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(ransomNote); i++ {
		note[ransomNote[i]]++
	}

	for j := 0; j < len(magazine); j++ {
		mag[magazine[j]]++
	}

	for key, value := range note {
		if value > mag[key] {
			return false
		} 
	}
	
	return true
}