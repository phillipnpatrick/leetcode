package Solutions

import (
	"fmt"
)

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:
// Input: nums = [3,2,3]
// Output: 3

// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

func majorityElement(nums []int) int {
	type Pair struct {
		key int
		value int
	}
	myMap := make(map[int]int)
	p := new(Pair)

	for _, value := range nums {
		myMap[value]++
	}

	for key, value := range myMap {
		fmt.Printf("Key %d, Value: %d\n", key, value)
		if p == nil {
			p.key = key
			p.value = value
		} else if value > p.value {
			p.key = key
			p.value = value
		}
	}

	return p.key
}