package Solutions

import (
	// "MyLeetCode/Solutions"
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/decode-string/description/?envType=study-plan-v2&envId=leetcode-75

// Given an encoded string, return its decoded string.
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
// Note that k is guaranteed to be a positive integer.
// You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc.
// Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
// For example, there will not be input like 3a or 2[4].
// The test cases are generated so that the length of the output will never exceed 105.

// Example 1:
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"

// Example 2:
// Input: s = "3[a2[c]]"
// Output: "accaccacc"

// Example 3:
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"

// "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
func decodeString(s string) string {
	for strings.ContainsAny(s, "[]") {
		str := getEncoded(s)
		old := str
		// fmt.Printf("str: %s \n", str)

		var decoded string
		if goodFormat(str){
			decoded = decode(str)
			str = cutString(str, str, decoded)
		} else {
			strToDecode, hasBrackets := getRightMostToDecode(str)

			for hasBrackets {
				decoded = decode(strToDecode)
				str = cutString(str, strToDecode, decoded)
				strToDecode, hasBrackets = getRightMostToDecode(str)
			}
		}		

		s = cutString(s, old, str)
		// fmt.Printf("s: %s \n", s)
	}

	return s
}

func getEncoded(s string) string {
	temp := s

	index := len(s) - 1
	count, left, right := -1, -1, -1

	for count != 0 {
		if s[index] == '[' {
			left = index
			count--
		} else if s[index] == ']' {
			if count < 0 {
				count = 1
				right = index + 1
			} else {
				count++
			}
		}
		index--
	}
	temp = s[:right]
	temp = temp[left:]

	if strings.Count(temp, "[") == 1 && strings.Count(temp, "]") == 1 {
		i := strings.Index(s, temp) - 1
		for i >= 0 {
			if isDigit(string(s[i])) {
				temp = fmt.Sprintf("%s%s", string(s[i]), temp)
			} else {
				break
			}
			i--
		}
	}

	return temp
}

func getRightMostToDecode(s string) (string, bool) {
	hasBrackets := true
	found := false
	temp := s

	for !found {
		if s[len(s)-1] == ']' {
			index := len(s) - 1
			left, right := -1, -1
			for left < 0 {
				if s[index] == '[' {
					left = index
				} else if s[index] == ']' {
					right = index
				}
				index--
			}
			temp = s[:left]
			j := -1
			for i := len(temp) - 1; i >= 0 && j < 0; i-- {
				if !isDigit(string(temp[i])) {
					j = i + 1
				}
			}
			temp = s[:right+1]
			if j == -1 {
				temp = strings.Replace(temp, "[", "", 1)
				temp = strings.Replace(temp, "]", "", 1)
				hasBrackets = false
			} else {
				temp = temp[j:]
			}
		} else {
			temp = temp[strings.LastIndex(temp, "]")+1:]
			hasBrackets = false
		}

		found = goodFormat(temp)
	}

	return temp, hasBrackets
}

func goodFormat(s string) bool {
	var num, inner string
	if strings.Contains(s, "[") {
		num = s[:strings.Index(s, "[")]
		inner = s[strings.Index(s, "["):]
	}

	if strings.Contains(s, "]") {
		inner = inner[:strings.LastIndex(inner, "]")+1]
	}

	if len(num) == 0 {
		return !strings.ContainsAny(inner, "[]")
	}

	return isDigit(num) && strings.Count(inner, "[") == 1 && strings.Count(inner, "]") == 1
}

func cutString(s string, old string, new string) string {
	if !strings.Contains(s, old) {
		left := strings.Index(old, "[")
		old = old[left+1:]
		right := strings.LastIndex(old, "]")
		old = old[:right]
	}

	s = strings.Replace(s, old, new, 1)

	return s
}

func isDigit(input string) bool {
	_, err := strconv.Atoi(input) // Try to convert the string to an integer
	return err == nil
}

func decode(s string) string {
	var sb strings.Builder
	var num int
	var encoded string

	for len(s) > 0 {
		leftBracket := strings.Index(s, "[")
		rightBracket := strings.Index(s, "]")

		if leftBracket < 0 && rightBracket < 0 {
			num = 1
			encoded = s
			s = ""
		} else {
			num, _ = strconv.Atoi(s[:leftBracket])
			encoded = s[leftBracket+1 : rightBracket]
			s = s[rightBracket+1:]
		}

		for i := 0; i < num; i++ {
			sb.WriteString(encoded)
		}
	}

	return sb.String()
}
