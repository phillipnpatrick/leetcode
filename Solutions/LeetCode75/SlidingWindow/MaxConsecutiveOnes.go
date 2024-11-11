package Solutions

import (
	"fmt"
	"math"
)

// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

// Example 1:
//                                    1
//                0 1 2 3 4 5 6 7 8 9 0
// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
//                                    1
//                0 1 2 3 4 5 6 7 8 9 0
// Explanation:  [1,1,1,0,0,1,1,1,1,1,1] -> nums[5:10]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// Example 2:
//                                    1 1 1 1 1 1 1 1 1
//                0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8
// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
//                                    1 1 1 1 1 1 1 1 1
//                0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8
// Explanation:  [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1] -> nums[2:11]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// [1,1,1,0,0,0,1,1,1,1,0,1], k=2; output: 7
func longestOnes(nums []int, k int) int {
	// return mySolution(nums, k)

	return fastSolution(nums, k)
}

func mySolution(nums []int, k int) int {
	max := float64(0)
	left := 0
	var zeros []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros = append(zeros, i)
		}
		if len(zeros) > k {
			fmt.Printf("len(nums[%2d:%2d]) = %d \t zeros: %v \n", left , i-1, len(nums[left:i]), zeros)
			max = math.Max(max, float64(len(nums[left:i])))
			left = zeros[0]+1
			zeros = zeros[1:]
		}
	}
	fmt.Printf("*len(nums[%2d:%2d]) = %d \t zeros: %v \n", left , len(nums)-1, len(nums[left:]), zeros)
	max = math.Max(max, float64(len(nums)-left))

	return int(max)
}

func fastSolution(nums []int, k int) int {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			k -= 1
		}

		if k < 0 {
			if nums[left] == 0 {
				k += 1
			}
			left += 1
		}
		right++
	}

	return right - left + 1
}
