package main

import (
	"MyLeetCode/Solutions"
	"fmt"
)

func main() {
    fmt.Println("Hello, world!")

	// input := "IceCreAm"
	// output := Solutions.ReverseVowels(input)

	// fmt.Printf("input: %s \t output: %s \n", input, output)

	// Solutions.RotateTheMatrix([][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}})
	// Solutions.RotateTheMatrix([][]int{{1,2,3},{4,5,6},{7,8,9}})
	// Solutions.RotateTheMatrix([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}})
	Solutions.RotateTheMatrix([][]int{{25,20,15,10,5},{24,19,14,9,4},{23,18,13,8,3},{22,17,12,7,2},{21,16,11,6,1}})
}
