package Solutions

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) 
// of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

// Example 1:
// Input: s = "abc", t = "ahbgdc"
// Output: true

// Example 2:
// Input: s = "axc", t = "ahbgdc"
// Output: false

// Input: s = "aec", t = "abcdefghijklmnop"
// Output: false

// Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence. 
// 			  In this scenario, how would you change your code?

// s = "abc", t = "zyxwvuahbgdc"				// true
// s = "abc", t = "zyxwvutsrqponmlkjihgfed"		// false
// s: "aaaaaa", t: "bbaaaa"						// true
func isSubsequence(s string, t string) bool {
	ptrS := 0
	sliceS := []rune(s)
	ptrT := 0
	sliceT := []rune(t)

	for ptrS < len(sliceS) {
		for ptrT < len(sliceT) && sliceS[ptrS] != sliceT[ptrT] {
			ptrT++
		}
		if ptrT >= len(sliceT) {
			break
		}
		sliceS = sliceS[ptrS+1:]
		ptrS = 0
		ptrT++
	}

	return len(sliceS) == 0
}

func isSubsequenceSimple(s string, t string) bool {
	k := 0
	for i := 0; i < len(t); i++ {
		if k < len(s) && s[k] == t[i] {
			k++
		}

	}
	return len(s) == k
}