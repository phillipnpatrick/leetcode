package Solutions

import (
	"fmt"
	"strings"
)

// Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.
// A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

// Example 2:
// Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
// Output: 1
// Explanation: There is 1 equal row and column pair:
// - (Row 2, Column 1): [2,7,7]

// Example 2:
// Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
// Output: 3
// Explanation: There are 3 equal row and column pairs:
// - (Row 0, Column 0): [3,1,2,2]
// - (Row 2, Column 2): [2,4,2,2]
// - (Row 3, Column 2): [2,4,2,2]

func equalPairs(grid [][]int) int {
	var rows, columns []string
	var sbRow,sbColumn strings.Builder
	pair := 0

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			sbRow.WriteString(fmt.Sprintf("%3d,", grid[row][column]))
			sbColumn.WriteString(fmt.Sprintf("%3d,", grid[column][row]))
		}
		rows = append(rows, sbRow.String())
		columns = append(columns, sbColumn.String())
		sbRow.Reset()
		sbColumn.Reset()
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(columns); j++ {
			if rows[i] == columns[j] {
				pair++
			}
		}
	}

	return pair
}

func itWorks(grid [][]int) int {
	var rows, columns []string
	var sb strings.Builder

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			sb.WriteString(fmt.Sprintf("%3d,", grid[row][column]))
		}
		rows = append(rows, sb.String())
		sb.Reset()
	}
	for column := 0; column < len(grid); column++ {
		for row := 0; row < len(grid[column]); row++ {
			sb.WriteString(fmt.Sprintf("%3d,", grid[row][column]))
		}
		columns = append(columns, sb.String())
		sb.Reset()
	}

	count := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(columns); j++ {
			if rows[i] == columns[j] {
				count++
			}
		}
	}

	return count
}
