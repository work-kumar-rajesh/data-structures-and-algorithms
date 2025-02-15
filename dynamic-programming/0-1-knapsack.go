package main

//similar to partition-equl-subset-sum
//maximize value with limit W

func knapSack(W int, weights, values []int) int {
	n := len(weights)
	dp := make([]int, W+1)

	for i := 0; i < n; i++ {
		for w := W; w >= weights[i]; w-- {
			dp[w] = max(dp[w], dp[w-weights[i]]+values[i])
		}
	}

	return dp[W]
}
