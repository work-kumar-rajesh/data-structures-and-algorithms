package main

import (
	"math"
	"sort"
)

// Given a wooden stick of length n units. The stick is labelled from 0 to n.
// For example, a stick of length 6 is labelled as follows:
// Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.

// You should perform the cuts in order, you can change the order of the cuts as you wish.

// The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts.
// When you cut a stick, it will be split into two smaller sticks (i.e. the sum of their lengths is the
// length of the stick before the cut). Please refer to the first example for a better explanation.

// Return the minimum total cost of the cuts.

func minCost(n int, cuts []int) int {
	m, cuts := len(cuts)+2, append(append([]int{0}, cuts...), n)
	sort.Ints(cuts)

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for length := 2; length < m; length++ {
		for i := 0; i < m-length; i++ {
			j := i + length
			dp[i][j] = cuts[j] - cuts[i]
			min_cost := math.MaxInt32
			for k := i + 1; k < j; k++ {
				min_cost = min(min_cost, dp[i][k]+dp[k][j])
			}
			dp[i][j] += min_cost
		}
	}

	return dp[0][m-1]
}
// dp[i][j] represents minimum cost to do all the cuts between i and j 