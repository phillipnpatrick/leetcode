package Solutions

import (
	"fmt"
)

// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, return the maximum number of points that lie on the same straight line.

// Example 1:
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 3

// Example 2:
// Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// Output: 4

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

// Line represents a line in general form: Ax + By + C = 0.
type Line struct {
	A, B, C int
	dx, dy, x1, y1 int
}

func (line Line) String() string {
	if line.dx == 0 {
		return fmt.Sprintf("x = %d", line.x1)
	}

	if line.dy == 0 {
		return fmt.Sprintf("y = %d", -line.y1)
	}

	// y - y1 = m(x - x1)
	return fmt.Sprintf("y - %d = (%d/%d)(x - %d)", line.y1, line.dy, line.dx, line.x1)
}

func maxPoints(points [][]int) int {
	// Map to store lines as keys and sets of unique points as values.
	lines := make(map[string]map[Point]struct{})

	for i := 0; i < len(points)-1; i++ {
		p := Point{x: points[i][0], y: points[i][1]}

		for j := i + 1; j < len(points); j++ {
			p1 := Point{x: points[j][0], y: points[j][1]}

			// Calculate the line for the two points
			line := equationOfLine(p, p1)

			// Initialize the map for this line if it doesn't exist
			if _, exists := lines[line.String()]; !exists {
				lines[line.String()] = make(map[Point]struct{})
			}

			// Add points to the line
			if isPointOnLine(p, line) {
				lines[line.String()][p] = struct{}{}
			}
			if isPointOnLine(p1, line) {
				lines[line.String()][p1] = struct{}{}
			}

			for k := j + 1; k < len(points); k++ {
				p2 := Point{x: points[k][0], y: points[k][1]}

				if isPointOnLine(p2, line) {
					lines[line.String()][p2] = struct{}{}
				}
			}
		}
	}

	count := 1

	// Print out each line and its unique points
	for _, pointsOnLine := range lines {
		if len(pointsOnLine) > count {
			count = len(pointsOnLine)
		}
	}

	return count
}

// normalizeLine generates a unique representation for a line using integer coefficients.
func normalizeLine(A, B, C int) Line {
	// Ensure that the sign of A is positive for consistency.
	if A < 0 || (A == 0 && B < 0) {
		A, B, C = -A, -B, -C
	}
	dx, dy, y1, x1 := 0, 0, 0, 0
	return Line{A, B, C, dx, dy, y1, x1}
}

func equationOfLine(p1, p2 Point) Line {
	dy := p2.y - p1.y
	dx := p2.x - p1.x
	A := dy
	B := dx
	C := -(dy*p1.x - dx*p1.y)

	line := normalizeLine(A, B, C)
	
	line.dy = dy
	line.dx = dx

	line.y1 = p1.y
	line.x1 = p1.x

	return line
}

func isPointOnLine(p Point, l Line) bool {
	if l.dx == 0 {
		// this is a vertical line or x = value
		// Line represents a line in general form: Ax + By + C = 0.
		
		return l.A*p.x + l.B*p.y == -l.C
	}

	// y - y1 = m(x - x1)
	left := float64(p.y - l.y1)
	right := float64(l.dy * (p.x - l.x1))/float64(l.dx)

	return left == right
}