package Solutions

import (
	"fmt"
	"math/big"
)

func trailingZeroes(n int) int {
	value := factorial(n)

	count := 0

	str := fmt.Sprintf("%d", value)

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '0' {
			count++
		} else {
			break
		}
	}

	return count
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= int64(n); i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}
