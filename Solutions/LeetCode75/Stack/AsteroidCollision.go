package Solutions

import "math"

// We are given an array asteroids of integers representing asteroids in a row.
// For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left).
// Each asteroid moves at the same speed.
// Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode.
// If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

// Example 1:
// Input: asteroids = [5,10,-5]
// Output: [5,10]
// Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

// Example 2:
// Input: asteroids = [8,-8]
// Output: []
// Explanation: The 8 and -8 collide exploding each other.

// Example 3:
// Input: asteroids = [10,2,-5]
// Output: [10]
// Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, asteroid := range asteroids {

		if len(stack) == 0 {
			stack = append(stack, asteroid)
		} else {
			top := stack[len(stack)-1]

			if top > 0 && asteroid > 0 { // if both asteroids are moving in the same direction they will never meet
				stack = append(stack, asteroid)
			} else if top < 0 && asteroid < 0 { // if both asteroids are moving in the same direction they will never meet
				stack = append(stack, asteroid)
			} else if top < 0 && asteroid > 0 { // asteroids are moving away from each other
				stack = append(stack, asteroid)
			} else {
				for i := len(stack) - 1; i >= 0; i-- {
					top = stack[i]

					if (top < 0 && asteroid < 0) || (top > 0 && asteroid > 0) {
						stack = append(stack, asteroid)
						break
					} else {
						absTop := int(math.Abs(float64(top)))
						absAsteroid := int(math.Abs(float64(asteroid)))

						if absTop > absAsteroid { // incoming asteroid destroyed
							break
						} else if absTop == absAsteroid { // asteroids were the same size, both explode
							stack = stack[:i]
							break
						}

						// incoming asteroid destroyed top asteroid
						stack = stack[:i]

						if len(stack) == 0 {
							stack = append(stack, asteroid)
						}
					}
				}
			}
		}
	}

	return stack
}
