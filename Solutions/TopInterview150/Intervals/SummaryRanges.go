package Solutions

import (
	"fmt"
	"strconv"
	"strings"
)

// You are given a sorted unique integer array nums.
// A range [a,b] is the set of all integers from a to b (inclusive).
// Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
// That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
// Each range [a,b] in the list should be output as:
// "a->b" if a != b
// "a" if a == b

// Example 1:
// Input: nums = [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
// Explanation: The ranges are:
// [0,2] --> "0->2"
// [4,5] --> "4->5"
// [7,7] --> "7"

// Example 2:
// Input: nums = [0,2,3,4,6,8,9]
// Output: ["0","2->4","6","8->9"]
// Explanation: The ranges are:
// [0,0] --> "0"
// [2,4] --> "2->4"
// [6,6] --> "6"
// [8,9] --> "8->9"

func summaryRanges(nums []int) []string {
	var working strings.Builder
	var output []string

	if len(nums) == 1 {
		output = append(output, strconv.Itoa(nums[0]))
	} else if len(nums) > 1 {
		working.WriteString(strconv.Itoa(nums[0]))
		
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] + 1 {
				if working.String() != strconv.Itoa(nums[i-1]) {
					working.WriteString(fmt.Sprintf("->%d", nums[i-1]))
				}
				output = append(output, working.String())
				working.Reset()
				working.WriteString(strconv.Itoa(nums[i]))
			} 
			
			if i == len(nums) - 1 {
				if working.String() != strconv.Itoa(nums[i]) {
					working.WriteString(fmt.Sprintf("->%d", nums[i]))
				}
				output = append(output, working.String())
			}
		}
	}

	return output
}
