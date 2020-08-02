// using top-down dp with extra space O(n)
// Time complexity: O(mn)
// Space complexity: O(mn)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    f := make([][]int, 2)
    for i:=0; i<len(obstacleGrid); i++ {
        f[0] = make([]int, len(obstacleGrid[i]))
        for j:=0; j<len(obstacleGrid[i]); j++ {
            if obstacleGrid[i][j] == 1 {
                f[0][j] = 0
                continue
            }
            if i==0 && j==0 {
                f[0][0] = 1
                continue
            }
            if i==0 {
                if obstacleGrid[0][j-1] == 1 {
                    f[0][j] = 0
                } else {
                    f[0][j] = f[0][j-1]
                }
            } else if j==0 {
                if obstacleGrid[1][j] == 1 {
                    f[0][j] = 0
                } else {
                    f[0][j] = f[1][j]
                }
            } else {
                if obstacleGrid[i-1][j] == 1 {
                    if obstacleGrid[i][j-1] == 1 {
                        f[0][j] = 0
                        continue
                    }
                    f[0][j] = f[0][j-1]
                }
                if obstacleGrid[i][j-1] == 1 {
                    f[0][j] = f[1][j]
                    continue
                }
                f[0][j] = f[1][j] + f[0][j-1]
            }
        }
        f[0], f[1] = f[1], f[0]
    }
    return f[1][len(obstacleGrid[0])-1]
}
