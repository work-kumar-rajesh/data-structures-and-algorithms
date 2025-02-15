package main

// Given an integer array nums, return true if you can partition the array into two
// subsets such that the sum of the elements in both subsets is equal or false otherwise.

// Example 1:

// Input: nums = [1,5,11,5]
// Output: true
// Explanation: The array can be partitioned as [1, 5, 5] and [11].

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[sum]
}

// Notice here the loop for total is inside ( lin no 25 ) and not outer as in coin change beacuse here each
//element can be used only once ( whereas in coin change each coin can be used infinite times )
