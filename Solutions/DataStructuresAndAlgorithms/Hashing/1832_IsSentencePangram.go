package Solutions

// https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/
// A pangram is a sentence where every letter of the English alphabet appears at least once.
// Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

// Example 1:
// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// Explanation: sentence contains at least one of every letter of the English alphabet.

// Example 2:
// Input: sentence = "leetcode"
// Output: false

func checkIfPangram(sentence string) bool {
    letters := make(map[byte]int)

	for i:=0; i < len(sentence); i++ {
		letters[sentence[i]]++
	}

	return len(letters) == len("abcdefghijklmnopqrstuvwxyz")
}