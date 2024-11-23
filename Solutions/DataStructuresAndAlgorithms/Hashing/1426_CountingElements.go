package Solutions

// https://leetcode.com/problems/counting-elements/description/
// Given an integer array arr, count how many elements x there are, such that x + 1 is also in arr. If there are duplicates in arr, count them separately.

// Example 1:
// Input: arr = [1,2,3] ... 1=>2, 2=>3
// Output: 2
// Explanation: 1 and 2 are counted cause 2 and 3 are in arr.

// Example 2:
// Input: arr = [1,1,3,3,5,5,7,7]
// Output: 0
// Explanation: No numbers are counted, cause there is no 2, 4, 6, or 8 in arr.

// Ex3 [1,3,2,3,5,0] --> 3 ... 0=>1, 1=>2, 2=>3
// Ex4 [1,1,2,2] --> 2 ... 1=>2, 1=>2
// Ex5 [1,1,2] --> 2
func countElements(arr []int) int {
    all := make(map[int]int)
	count := 0

	for i := 0; i < len(arr); i++ {
		all[arr[i]]++
	}

	for i := 0; i < len(arr); i++ {
		if all[arr[i]+1] > 0 {
			count++
		}
	}

	return count
}