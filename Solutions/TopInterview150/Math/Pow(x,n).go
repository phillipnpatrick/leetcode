package Solutions

import (
	"fmt"
	"math"
	"strconv"
)

// Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

// Example 1:
// Input: x = 2.00000, n = 10
// Output: 1024.00000

// Example 2:
// Input: x = 2.10000, n = 3
// Output: 9.26100

// Example 3:
// Input: x = 2.00000, n = -2
// Output: 0.25000
// Explanation: 2^(-2) = (1/2)^2 = 1/4 = 0.25

func myPow(x float64, n int) float64 {
	var product float64

	if n == 0 || x == 1{
		product = 1
	} else if x == -1 {
		pow := int(math.Abs(float64(n)))

		if pow%2 == 0 {
			product = 1
		} else {
			product = -1
		}
	} else {
		product = x
		pow := int(math.Abs(float64(n)))

		for i := 2; i <= pow; i++ {
			product *= x
			
			if isOutOfRange(product) || (n < 0 && isOutOfRange(1/product)) {
				break
			}
		}

		// negative exponent
		if n < 0 {
			product = 1 / product
		}

		product, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", product), 64)
	}

	return product
}

func isOutOfRange(num float64) bool {
	return num > 0 && num < 0.000001
}