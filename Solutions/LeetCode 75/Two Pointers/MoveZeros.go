package Solutions

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

// Hint 1: In-place means we should not be allocating any space for extra array. But we are allowed to modify the existing array.
//         However, as a first step, try coming up with a solution that makes use of additional space.
//         For this problem as well, first apply the idea discussed using an additional array and the in-place solution will pop up eventually.

// Hint 2: A two-pointer approach could be helpful here. The idea would be to have one pointer for iterating the array and another pointer that
//		   just works on the non-zero elements of the array
func moveZeroes(nums []int) []int {
	lastNonZeroFoundAt := 0
    
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
            lastNonZeroFoundAt++
        }
    }

	// ptrZero := 0
	// ptrNonZero := 0

	// for ptrZero < len(nums) && ptrNonZero < len(nums) {

	// 	// read until array value is zero
	// 	for ptrZero < len(nums) && nums[ptrZero] != 0 {
	// 		ptrZero++
	// 	}

	// 	ptrNonZero = ptrZero

	// 	// read until array value is non-zero
	// 	for ptrNonZero < len(nums) && nums[ptrNonZero] == 0 {
	// 		ptrNonZero++
	// 	}

	// 	if ptrZero < len(nums) && ptrNonZero < len(nums) {
	// 		fmt.Printf("nums[%v]: %v, nums[%v]: %v\n", ptrZero, nums[ptrZero], ptrNonZero, nums[ptrNonZero])

	// 		// swap non-zero with zero
	// 		temp := nums[ptrZero]
	// 		nums[ptrZero] = nums[ptrNonZero]
	// 		nums[ptrNonZero] = temp
	// 	}

	// 	ptrZero++
	// }

	return nums
}
