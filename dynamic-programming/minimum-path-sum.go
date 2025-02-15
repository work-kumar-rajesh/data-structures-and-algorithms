package main

// Given a m x n grid filled with non-negative numbers, find a path from top
//  left to bottom right, which minimizes the sum of all numbers along its path.

// Note: You can only move either down or right at any point in time.

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[m-1][n-1]
}
