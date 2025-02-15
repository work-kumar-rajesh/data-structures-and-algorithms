package main

import "sort"

// Given an integer array nums, return the length of the longest strictly increasing
// subsequence
// Example 1:
// Input: nums = [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

//N^2 dp solution

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

//Nlogn approach

// The algorithm maintains a helper slice sub where each element represents the smallest possible tail of an increasing subsequence of a given length.
// If the current number is greater than the last element of sub, it extends the subsequence; otherwise,
// it finds the position to replace using binary search.
// Replacing an element with a smaller one keeps the potential subsequence "open" for future extensions by lowering the tail value.
// Although sub doesn't always form a valid subsequence, its length always equals the length of the longest increasing subsequence found so far.
// Using binary search for replacements ensures an efficient O(n log n) time complexity.

func lengthOfLIS2(nums []int) int {
	sub := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > sub[len(sub)-1] {
			sub = append(sub, num)
		} else {
			j := sort.Search(len(sub), func(m int) bool { return sub[m] >= num })
			sub[j] = num
		}
	}
	return len(sub)
}
