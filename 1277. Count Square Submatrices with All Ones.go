// Time complexity: O(mn)
// Space complexity: O(n)

func countSquares(matrix [][]int) int {
    if len(matrix) == 0 {
        return 0
    }
    result := 0
    dp := make([][]int, 2)
    for i:=0; i<len(matrix); i++ {
        dp[0] = make([]int, len(matrix[i]))
        for j:=0; j<len(matrix[i]); j++ {
            if i==0 || j==0 {
                dp[0][j] = matrix[i][j]
            } else if matrix[i][j] == 1 {
                dp[0][j] = min(dp[1][j-1], dp[1][j], dp[0][j-1]) + 1
            }
            result += dp[0][j]
        }
        dp[0], dp[1] = dp[1], dp[0]
    }
    return result
}

func min(x, y, z int) int {
    if x<=y && x<=z {
        return x
    } else if y<=x && y<=z {
        return y
    }
    return z
}