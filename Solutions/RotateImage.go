package Solutions

import "fmt"

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]

// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

func RotateTheMatrix(matrix [][]int) {
	rotate(matrix)
}

func transpose(matrix [][]int) {
	// swap row and column indices
	for row := 0; row < len(matrix); row++ {
		for column := row; column < len(matrix[row]); column++ {
			matrix[row][column], matrix[column][row] = matrix[column][row], matrix[row][column]
		}
	}
}

func reverseRows(matrix [][]int) {
	upperBound := len(matrix) - 1

	for row := 0; row < len(matrix); row++ {
		for columnLeft, columnRight := 0, upperBound; columnLeft < columnRight; columnLeft, columnRight = columnLeft+1, columnRight-1 {
			matrix[row][columnLeft], matrix[row][columnRight] = matrix[row][columnRight], matrix[row][columnLeft]
		}
	}
}

func rotate(matrix [][]int) [][]int {	
	transpose(matrix)
	print(matrix)
	reverseRows(matrix)
	print(matrix)

	return matrix
}

func print(matrix [][]int) {
	for i:=0; i < len(matrix); i++ {
		for j:=0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}