package Solutions

import "slices"

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
	runeRansomNote := []rune(ransomNote)
	runeMagazine := []rune(magazine)

	slices.Sort(runeRansomNote)
	slices.Sort(runeMagazine)

	idxRansomNote, idxMagazine := 0, 0

	for idxRansomNote < len(runeRansomNote) && idxMagazine < len(runeMagazine) {
		if runeRansomNote[idxRansomNote] == runeMagazine[idxMagazine] {
			idxRansomNote++
		}
		idxMagazine++
	}

	return idxRansomNote == len(ransomNote)
}