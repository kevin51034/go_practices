// using top-down dp with extra space O(n)
// Time complexity: O(mn)
// Space complexity: O(mn)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    d := make([][]int, 2)
    for i:=0; i<len(obstacleGrid); i++ {
        d[0] = make([]int, len(obstacleGrid[i]))
        for j:=0; j<len(obstacleGrid[i]); j++ {
            if obstacleGrid[i][j] == 1 {
                d[0][j] = 0
                continue
            }
            if i==0 && j==0 {
                d[0][j] = 1
                continue
            }
            if i==0 {
                d[0][j] = d[0][j-1]
            } else if j==0 {
                d[0][j] = d[1][j]
            } else {
                d[0][j] = d[1][j] + d[0][j-1]
            }
        }
        d[0], d[1] = d[1], d[0]
    }
    return d[1][len(d[1])-1]
}
