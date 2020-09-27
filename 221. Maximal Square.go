// Time complexity: O(mn)
// Space complexity: O(mn)

func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    result := 0
    dp := make([][]int, len(matrix))
    for i:=0; i<len(matrix); i++ {
        dp[i] = make([]int, len(matrix[i]))
        for j:=0; j<len(matrix[i]); j++ {
            if i==0 || j==0 {
                dp[i][j] = int(matrix[i][j] - '0')
            } else if matrix[i][j] == '1' {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
            }
            if dp[i][j] > result {
                result = dp[i][j]
            }
        }
    }
    return result*result
}

func min(x, y, z int) int {
    if x<=y && x<=z {
        return x
    } else if y<=x && y<=z {
        return y
    }
    return z
}

// Time complexity: O(mn)
// Space complexity: O(n)

func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    result := 0
    dp := make([][]int, 2)
    for i:=0; i<len(matrix); i++ {
        dp[0] = make([]int, len(matrix[i]))
        for j:=0; j<len(matrix[i]); j++ {
            if i==0 || j==0 {
                dp[0][j] = int(matrix[i][j] - '0')
            } else if matrix[i][j] == '1' {
                dp[0][j] = min(dp[1][j-1], dp[1][j], dp[0][j-1]) + 1
            }
            if dp[0][j] > result {
                result = dp[0][j]
            }
        }
        dp[0], dp[1] = dp[1], dp[0]
    }
    return result*result
}

func min(x, y, z int) int {
    if x<=y && x<=z {
        return x
    } else if y<=x && y<=z {
        return y
    }
    return z
}