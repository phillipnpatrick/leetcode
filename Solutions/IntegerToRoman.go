package Solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	roman := ""
	symbols := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	subtractiveForm := map[int]string{
		4:   "IV",
		9:   "IX",
		40:  "XL",
		90:  "XC",
		400: "CD",
		900: "CM",
	}

	for q := len(strconv.Itoa(num)) - 1; num > 0; q-- {
		u := int(math.Pow10(q))
		p := u * (num / u)

		num -= p
		temp := ""

		if strings.HasPrefix(strconv.Itoa(p), "4") || strings.HasPrefix(strconv.Itoa(p), "9") {
			temp = subtractiveForm[p]
		} else if value, exists := symbols[p]; exists {
			temp = value
		} else {
			for p > 0 {
				symbolIndex := closestValue(p)
				temp = fmt.Sprintf("%s%s", temp, symbols[symbolIndex])
				p -= symbolIndex
			}
		}

		roman = fmt.Sprintf("%s%s", roman, temp)
	}

	return roman
}

func closestValue(target int) int {
	values := []int{1, 5, 10, 50, 100, 500, 1000}
	closest := values[0]
	minDiff := math.Abs(float64(closest - target))

	for _, value := range values[1:] {
		diff := math.Abs(float64(value - target))
		if diff < minDiff && value <= target {
			minDiff = diff
			closest = value
		}
	}
	return closest
}
