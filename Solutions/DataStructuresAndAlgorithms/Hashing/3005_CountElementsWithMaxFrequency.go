package Solutions

// https://leetcode.com/problems/count-elements-with-maximum-frequency/description/
// You are given an array nums consisting of positive integers.
// Return the total frequencies of elements in nums such that those elements all have the maximum frequency.
// The frequency of an element is the number of occurrences of that element in the array.

// Example 1:
// Input: nums = [1,2,2,3,1,4]
// Output: 4
// Explanation: The elements 1 and 2 have a frequency of 2 which is the maximum frequency in the array.
// So the number of elements in the array with maximum frequency is 4.

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: 5
// Explanation: All elements of the array have a frequency of 1 which is the maximum.
// So the number of elements in the array with maximum frequency is 5.

func maxFrequencyElements(nums []int) int {
	frequencies := make(map[int]int)
    max := -1
	elements := 0

	for i := 0; i < len(nums); i++ {
		frequencies[nums[i]]++
		max = getMax(max, frequencies[nums[i]])
	}

	for _, value := range frequencies {
		if value == max {
			elements += value
		}
	}

	return elements
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}