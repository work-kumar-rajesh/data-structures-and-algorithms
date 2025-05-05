```go
func numTilings(n int) int {
    if n <= 2 { return n }
    if n == 3 { return 5 }
    var mod int = 1e9+7
    dp := make([]int,n+1)
    dp[1],dp[2],dp[3] = 1,2,5
    for i := 4 ; i <= n ; i++ {
        dp[i] = (2*dp[i-1]+dp[i-3])%mod
    }
    return dp[n]
}
```