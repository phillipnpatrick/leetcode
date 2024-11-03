package Solutions

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Example 2:
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Hint 1: Think how you can efficiently utilize prefix and suffix products to calculate the product of all elements except self for each index. 
// 		   Can you pre-compute the prefix and suffix products in linear time to avoid redundant calculations?

// Hint 2: Can you minimize additional space usage by reusing memory or modifying the input array to store intermediate results?

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)

	leftProduct := 1
	rightProduct := 1
	prefix[0] = 1
	suffix[n-1] = 1

	for i := 1; i < n; i++ {
		leftProduct *= nums[i-1]
		rightProduct *= nums[n-i]

		prefix[i] = leftProduct
		suffix[(n-1)-i] = rightProduct
	}

	answer := make([]int, n)

	for j := 0; j < n; j++ {
		answer[j] = prefix[j] * suffix[j]
	}

    return answer
}