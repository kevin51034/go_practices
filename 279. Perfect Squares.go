// Time complexity: O(n * sqrt(n))
// Space complexity: O(n)

func numSquares(n int) int {
    dp := make([]int, n+1)
    for i:=0; i<len(dp); i++ {
        dp[i] = math.MaxInt64
    }
    dp[0] = 0
    for i:=1; i<=n; i++ {
        for j:=1; j*j<=i; j++ {
            dp[i] = min(dp[i], dp[i-j*j]+1)
        }
    }
    return dp[n]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}