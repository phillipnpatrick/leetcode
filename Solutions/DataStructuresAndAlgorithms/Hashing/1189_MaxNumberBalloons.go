package Solutions

import "math"

// https://leetcode.com/problems/maximum-number-of-balloons/description/
// Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.
// You can use each character in text at most once. Return the maximum number of instances that can be formed.

// Example 1:
// Input: text = "nlaebolko"
// Output: 1

// Example 2:
// Input: text = "loonbalxballpoon"
// Output: 2

// Example 3:
// Input: text = "leetcode"
// Output: 0

func maxNumberOfBalloons(text string) int {
	chars := make(map[byte]int)
	balloons := 0

	for i := 0; i < len(text); i++ {
		letter := text[i]
		if letter == 'b' || letter == 'a' || letter == 'l' || letter == 'o' || letter == 'n' {
			chars[letter]++
		}
	}

	loMin := smallest(chars['l'], chars['o'])
	loMin /= 2
	
	abnMin := smallest(chars['a'],chars['b'],chars['n'])
	
	balloons = smallest(loMin, abnMin)

	return balloons
}

func smallest(values ...int) int {
	lowest := math.MaxInt

	for _, value := range values {
		if value < lowest {
			lowest = value
		}
	}

	return lowest
}
