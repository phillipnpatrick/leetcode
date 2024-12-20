package Solutions

// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
// You must not use any built-in exponent function or operator.
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

// Example 1:
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.

// Example 2:
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

func mySqrt(x int) int {
	sqrt := -1

	for y := 0; sqrt == -1; y++ {
		if x == y*y {
			sqrt = y
		} else if x > y*y && x < (y+1)*(y+1) {
			sqrt = y
		}
	}

	return sqrt
}
