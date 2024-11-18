package Solutions

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/counting-bits/description/?envType=study-plan-v2&envId=leetcode-75

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

// Example 1:
// Input: n = 2
// Output: [0,1,1]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10

// Example 2:
// Input: n = 5
// Output: [0,1,1,2,1,2]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101

func countBits(n int) []int {
	result := []int{}

	for i := int64(0); i <= int64(n); i++ {
		binary := strconv.FormatInt(i, 2)
	
		ones := strings.Count(binary, "1")

		result = append(result, ones)
	}

    return result
}