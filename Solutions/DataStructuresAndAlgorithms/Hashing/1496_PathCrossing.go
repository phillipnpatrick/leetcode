package Solutions

// https://leetcode.com/problems/path-crossing/description/
// Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively.
// You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.

// Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited. Return false otherwise.

// Example 1:
// Input: path = "NES"
// Output: false
// Explanation: Notice that the path doesn't cross any point more than once.

// Example 2:
// Input: path = "NESWW"
// Output: true
// Explanation: Notice that the path visits the origin twice.

type Coordinate struct {
	x, y int
}

func isPathCrossing(path string) bool {
	myPath := make(map[Coordinate]int)
	curr := Coordinate{x: 0, y: 0}

	myPath[curr]++

	for i := 0; i < len(path); i++ {
		if path[i] == 'N' {
			curr.y++
		} else if path[i] == 'S' {
			curr.y--
		} else if path[i] == 'E' {
			curr.x++
		} else if path[i] == 'W' {
			curr.x--
		}
		myPath[curr]++

		if myPath[curr] > 1 {
			return true
		}
	}

	return false
}
