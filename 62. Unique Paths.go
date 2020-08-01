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
    f := make([][]int, 2)
    for i:=0; i<n; i++ {
        f[0] = make([]int, m)
        for j:=0; j<m; j++ {
            if i==0 && j==0 {
                f[0][j] = 1
                continue
            }
            if i==0 {
                f[0][j] = f[0][j-1]
            } else if j==0 {
                f[0][j] = f[1][j]
            } else {
                f[0][j] = f[1][j] + f[0][j-1]
            }
        }
        f[0], f[1] = f[1], f[0]
    }
    return f[1][m-1]
}