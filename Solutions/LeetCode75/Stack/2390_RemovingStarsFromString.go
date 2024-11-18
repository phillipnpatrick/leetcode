package Solutions

import (
	"MyLeetCode/Solutions"
	"fmt"
	"strings"
)

// https://leetcode.com/problems/removing-stars-from-a-string/description/?envType=study-plan-v2&envId=leetcode-75

// You are given a string s, which contains stars *.
// In one operation, you can:
// Choose a star in s.
// Remove the closest non-star character to its left, as well as remove the star itself.
// Return the string after all stars have been removed.
// Note:
// The input will be generated such that the operation is always possible.
// It can be shown that the resulting string will always be unique.

// Example 1:
// Input: s = "leet**cod*e"
// Output: "lecoe"
// Explanation: Performing the removals from left to right:
// - The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
// - The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
// - The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
// There are no more stars, so we return "lecoe".

// Example 2:
// Input: s = "erase*****"
// Output: ""
// Explanation: The entire string is removed, so we return an empty string.

func removeStars(s string) string {
	var sb strings.Builder
	var output []string
	count := 0
	index := len(s)-1

	for index >= 0 {
		value := string(s[index])
		if value == "*" {
			count++
			index--
		} else if count > 0 {
			for count > 0 && string(s[index]) != "*" {
				index--
				count--
			}
		} else {
			output = append([]string{value}, output...)
			index--
		}
	}

	for i:=0;i< len(output); i++{
		sb.WriteString(fmt.Sprintf("%v", output[i]))
	}

	return sb.String()
	// temp := fmt.Sprintf("%v", output)
	// temp = strings.ReplaceAll(temp, "[", "")
	// temp = strings.ReplaceAll(temp, " ", "")
	// temp = strings.ReplaceAll(temp, "]", "")
	// return temp
}

func withStack(s string) string {
	stack := Solutions.Stack[string]{}

	for i := 0; i < len(s); i++ {
		value := string(s[i])
		if value == "*" {
			if !stack.IsEmpty() {
				stack.Pop()
			}
		} else {
			stack.Push(value)
		}
	}

	return strings.ReplaceAll(stack.String(), ", ", "")
}
