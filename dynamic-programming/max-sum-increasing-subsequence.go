package main

// Given an array of positive integers arr. Find the maximum sum subsequence of the given array such that
// the integers in the subsequence are sorted in strictly increasing order i.e. a strictly increasing

func maxSumIS(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = arr[i]
	}
	ans := arr[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+arr[i])
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
