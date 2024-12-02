package Solutions

import "MyLeetCode/Solutions"

// https://leetcode.com/problems/valid-parentheses/description/
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// Example 4:
// Input: s = "([])"
// Output: true

func isValid(s string) bool {
	matching := make(map[rune]rune)
	matching['(']=')'
	matching['[']=']'
	matching['{']='}'

	stack := Solutions.Stack[rune]{}
    
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if _, exists := matching[c]; exists {
			stack.Push(c)
		} else {
			if stack.IsEmpty() {
				return false
			}

			previousOpening, _ := stack.Pop()
			if matching[previousOpening] != c {
				return false
			}
		}
	}

	return stack.IsEmpty()
}