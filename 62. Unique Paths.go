// solution 1
// using top-down dp with extra space O(n^2)
// Time complexity: O(mn)
// Space complexity: O(mn)

func uniquePaths(m int, n int) int {
    f := make([][]int, n)
    for i:=0; i<n; i++ {
        f[i] = make([]int, m)
        for j:=0; j<m; j++ {
            if i==0 && j==0 {
                f[i][j] = 1
                continue
            }
            if i==0 {
                f[i][j] = f[i][j-1]
            } else if j==0 {
                f[i][j] = f[i-1][j]
            } else {
                f[i][j] = f[i-1][j] + f[i][j-1]
            }
        }
    }
    return f[n-1][m-1]
}


// solution 2
// using top-down dp with extra space O(n)
// Time complexity: O(mn)
// Space complexity: O(mn) // because original space is O(mn)

func uniquePaths(m int, n int) int {
    dp := make([][]int, 2)
    for i:=0; i<m; i++ {
        dp[0] = make([]int, n)
        for j:=0; j<n; j++ {
            if i==0 && j==0 {
                dp[0][j] = 1
                continue
            }
            if i == 0 || j == 0 {
                dp[0][j] = 1
            } else {
                dp[0][j] = dp[1][j] + dp[0][j-1]
            }
        }
        dp[0], dp[1] = dp[1], dp[0]
    }
    return dp[1][n-1]
}