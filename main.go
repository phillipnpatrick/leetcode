package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
    fmt.Println("Hello, world!")

	// Example float64 number
	num := 123.456789

	intPart, fracPart := math.Modf(num)

	fmt.Printf("intPart: %f, fracPart: %f \n", intPart, fracPart)
	myModf(123.456789)

	myModf(2.00000)

	x := 0.00001
	y := 0.00009

	fmt.Printf("%v\n", x < y)
}

func myModf(num float64) (int, int, int) {
	n := fmt.Sprintf("%f", num)
	p := strings.Index(n, ".")
	
	i := string(n[:p])
	f := string(n[p+1:])
	
	count := len(f)

	integer, _ := strconv.Atoi(i)
	fraction, _ := strconv.Atoi(f)

	if fraction == 0 {
		count = fraction
	}

	fmt.Printf("intPart: %d \t fracPart: %d \t count: %d \n", integer, fraction, count)

	return integer, fraction, count
}
