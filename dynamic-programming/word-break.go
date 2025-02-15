package main

//partition dp

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence
//  of one or more dictionary words.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.

// Example 1:

// Input: s = "leetcode", wordDict = ["leet","code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = dict[s[i:i+1]]
	}

	for L := 2; L <= n; L++ {
		for i := 0; i <= n-L; i++ {
			j := i + L - 1
			if dict[s[i:j+1]] {
				dp[i][j] = true
				continue
			}
			for k := i; k < j; k++ {
				if dp[i][k] && dp[k+1][j] {
					dp[i][j] = true
					break
				}
			}
		}
	}

	return dp[0][n-1]
}
