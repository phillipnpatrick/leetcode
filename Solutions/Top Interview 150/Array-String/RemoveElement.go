package Solutions

// import (
// 	"math"
// )

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
// Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.

// Example 1:
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 2.
// It does not matter what you leave beyond the returned k (hence they are underscores).

// Example 2:
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
// Note that the five elements can be returned in any order.
// It does not matter what you leave beyond the returned k (hence they are underscores).

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index;
// 	count := 0
// 	left := 0
// 	right := len(nums) - 1

// 	for left < len(nums) && right >= 0 {
// 		for right >= 0 && nums[right] == val {
// 			count++
// 			nums[right] = math.MinInt8
// 			right--
// 		}
// 		for left < len(nums) - 1 && nums[left] != val {
// 			left++
// 		}
// 		if left < right {
// 			count++
// 			nums[left] = math.MinInt8
// 			nums[left], nums[right] = nums[right], nums[left]
// 			left++
// 			right--
// 		} else {
// 			break
// 		}
// 	}

// 	return len(nums) - count
}
