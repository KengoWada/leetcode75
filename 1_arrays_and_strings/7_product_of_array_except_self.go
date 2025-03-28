package arraysandstrings

func ProductExceptSelf(nums []int) []int {
	// Solution works but exceeds time limits
	// products := make([]int, 0, len(nums))
	// for i := range len(nums) {
	// 	sublist := slices.Delete(nums, i, i+1)
	// 	product := 1
	// 	for _, v := range sublist {
	// 		if v == 0 {
	// 			product = 0
	// 			break
	// 		}

	// 		product *= v
	// 	}
	// 	products[i] = product
	// }

	// return products

	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	ans := make([]int, n)
	left[0] = 1
	right[n-1] = 1

	for i := 1; i < n; i++ {
		j := n - i - 1
		left[i] = nums[i-1] * left[i-1]
		right[j] = nums[j+1] * right[j+1]
	}

	for i := range n {
		ans[i] = left[i] * right[i]
	}

	return ans
}
