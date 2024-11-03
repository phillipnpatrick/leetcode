package Solutions

func intToRoman(num int) string {
	n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	s := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	i := 0
	str := ""
	for num > 0 {
		if num >= n[i] {
			str = str + s[i]
			num -= n[i]
		} else {
			i++
		}
	}
	return str
}
