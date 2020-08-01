// solution 1
// using top-down dp with extra space O(n^2)
// Time complexity: O(mn)
// Space complexity: O(mn)

func minPathSum(grid [][]int) int {
    f := make([][]int, len(grid))
    for i:=0; i<len(grid); i++ {
        f[i] = make([]int, len(grid[i]))
        for j:=0; j<len(grid[i]); j++ {
            f[i][j] = grid[i][j]
            if i==0 && j==0 {
                continue
            }
            if i==0 {
                f[i][j] = f[i][j-1] + grid[i][j]
            } else if j==0 {
                f[i][j] = f[i-1][j] + grid[i][j]
            } else {
                f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
            }
        }
    }
    return f[len(grid)-1][len(f[len(grid)-1])-1]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}


// solution 2
// using top-down dp with extra space O(n)
// Time complexity: O(mn)
// Space complexity: O(mn) // because original space is O(mn)

func minPathSum(grid [][]int) int {
    f := make([][]int, 2)
    for i:=0; i<len(grid); i++ {
        f[0] = make([]int, len(grid[i]))
        for j:=0; j<len(grid[i]); j++ {
            f[0][j] = grid[i][j]
            if i==0 && j==0 {
                continue
            }
            if i==0 {
                f[0][j] = f[0][j-1] + grid[i][j]
            } else if j==0 {
                f[0][j] = f[1][j] + grid[i][j]
            } else {
                f[0][j] = min(f[1][j], f[0][j-1]) + grid[i][j]
            }
        }
        f[0], f[1] = f[1], f[0]
    }
    return f[1][len(f[1])-1]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}


// solution 3
// using top-down dp with extra space O(1)
// Time complexity: O(mn)
// Space complexity: O(mn) // because original space is O(mn)

func minPathSum(grid [][]int) int {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            if i==0 && j==0 {
                continue
            }
            if i==0 {
                grid[i][j] = grid[i][j-1] + grid[i][j]
            } else if j==0 {
                grid[i][j] = grid[i-1][j] + grid[i][j]
            } else {
                grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
            }
        }
    }
    return grid[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

