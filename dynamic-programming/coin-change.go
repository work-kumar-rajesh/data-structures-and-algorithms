package main

// You are given an integer array coins representing coins of different denominations and an integer amount
// representing a total amount of money.
// Return the fewest number of coins that you need to make up that amount. If that amount of
//  money cannot be made up by any combination of the coins, return -1.
// You may assume that you have an infinite number of each kind of coin.

// Example 1:

// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = 1e9
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		if coins[i] > amount {
			continue
		}
		dp[coins[i]] = 1
	}

	for i := 0; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == 1e9 {
		return -1
	}
	return dp[amount]
}
