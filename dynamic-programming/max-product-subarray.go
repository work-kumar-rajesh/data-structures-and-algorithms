package main

// Given an integer array nums, find a
// subarray
//  that has the largest product, and return the product.

// The test cases are generated so that the answer will fit in a 32-bit integer.

func maxProduct(nums []int) int {
	ans, imax, imin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imin = min(nums[i], nums[i]*imin)
		imax = max(nums[i], nums[i]*imax)
		ans = max(imax, ans)
	}
	return ans
}
