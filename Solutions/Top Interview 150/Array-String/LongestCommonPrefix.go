package Solutions

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var common []rune

	for i := 1; i < len(strs); i++ {
		var working []rune
		previous := []rune(strs[i-1])
		current := []rune(strs[i])
		
		for j := 0; j < len(previous) && j < len(current) && previous[j] == current[j]; j++ {
			working = append(working, previous[j])
		}

		if len(common) == 0 || len(working) < len(common) {
			common = working
		}

		if len(common) == 0 {
			break
		}
	}

	return string(common)
}