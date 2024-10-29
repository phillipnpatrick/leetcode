package Solutions

import "math"

// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that
// i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

func increasingTriplet(nums []int) bool {
	triplet1, triplet2 := math.MaxInt32, math.MaxInt32

    for i := 0; i < len(nums); i++ {
        if nums[i] <= triplet1 {
            triplet1 = nums[i]
        } else if nums[i] <= triplet2 {
            triplet2 = nums[i]
        } else {
            return true
        }
    }
    return false
}