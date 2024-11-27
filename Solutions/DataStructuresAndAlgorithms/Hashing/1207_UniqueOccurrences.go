package Solutions

// https://leetcode.com/problems/unique-number-of-occurrences/description/
// Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

// Example 1:
// Input: arr = [1,2,2,1,1,3]
// Output: true
// Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

// Example 2:
// Input: arr = [1,2]
// Output: false

// Example 3:
// Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
// Output: true

func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	set := make(map[int]struct{})

	for i := 0; i < len(arr); i++ {
		occurrences[arr[i]]++
	}

	for _, value := range occurrences {
		if _, exists := set[value]; exists {
			return false
		}
		set[value] = struct{}{}
	}

	return true
}
