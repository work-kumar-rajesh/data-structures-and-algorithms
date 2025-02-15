package main


// Given a string s, a partitioning of the string is a palindrome partitioning if every sub-string of the partition
//  is a palindrome. Determine the fewest cuts needed for palindrome partitioning of the given string.

// Examples:

// Input: s = "ababbbabbababa"
// Output: 3
// Explaination: After 3 partitioning substrings 
// are "a", "babbbab", "b", "ababa".


func palindromicPartition(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	isPal := make([][]bool, n)
	for i := range isPal {
		isPal[i] = make([]bool, n)
		isPal[i][i] = true 
	}
	for L := 2; L <= n; L++ {
		for i := 0; i <= n-L; i++ {
			j := i + L - 1
			if s[i] == s[j] {
				if L == 2 {
					isPal[i][j] = true
				} else {
					isPal[i][j] = isPal[i+1][j-1]
				}
			} else {
				isPal[i][j] = false
			}
		}
	}

	// dp[i] will hold the minimum cuts needed for substring s[0...i].
	// If s[0...i] is a palindrome, dp[i] = 0.
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if isPal[0][i] {
			dp[i] = 0
		} else {
			dp[i] = i // maximum cuts: cutting between every character.
			for j := 0; j < i; j++ {
				if isPal[j+1][i] && dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[n-1]
}
