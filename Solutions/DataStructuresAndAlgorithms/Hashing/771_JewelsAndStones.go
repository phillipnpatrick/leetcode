package Solutions

// https://leetcode.com/problems/jewels-and-stones/description/
// You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have.
// Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

// Letters are case sensitive, so "a" is considered a different type of stone from "A".

// Example 1:
// Input: jewels = "aA", stones = "aAAbbbb"
// Output: 3

// Example 2:
// Input: jewels = "z", stones = "ZZ"
// Output: 0

func numJewelsInStones(jewels string, stones string) int {
	mapJewels := make(map[byte]int)
	count := 0

	for i := 0; i < len(jewels); i++ {
		mapJewels[jewels[i]]++
	}

	for j := 0; j < len(stones); j++ {
		if _, exists := mapJewels[stones[j]]; exists {
			count++
		}
	}

	return count
}
