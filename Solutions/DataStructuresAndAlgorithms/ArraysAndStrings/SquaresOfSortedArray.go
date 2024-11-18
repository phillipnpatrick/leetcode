package Solutions

// nums = [-4,-1,0,3,10] ==> [0,1,9,16,100]
// nums = [-7,-3,2,3,11] ==> [4,9,9,49,121]
// nums: []int{-4,-1,0,3,5,6,7,8,9,10}  ==> [0,1,9,16,25,36,49,64,81,100]
// nums: []int{-5,-4,-3,-2,-1,0,1}, ==> []int{0,1,1,4,9,16,25},
func sortedSquares(nums []int) []int {
    ans := []int{}
    a := 0
    b := 1
    
    for a >= 0 && b < len(nums) {
        n1 := nums[a]
        n2 := nums[b]
        
        if n1 < 0 && n2 < 0 {
			a++
            b++
        } else if n1 < 0 && n2 == 0 {
            ans = append(ans, 0)
			b++
        } else if n1 < 0 && n2 > 0 {
            n1 = -1*n1
            
			if n1 < n2 {
				ans = append(ans, n1*n1)
				a--
			} else {
				ans = append(ans, n2*n2)
				b++
			}
        } else if n1 >= 0 && n2 >= 0 {
			b = a
			a = -1
			break
		}
    }

	for a >= 0 {
		ans = append(ans, nums[a] * nums[a])
		a--
	}

	for b < len(nums) {
		ans = append(ans, nums[b] * nums[b])
		b++
	}
    
    return ans
}