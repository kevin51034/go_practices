func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    result := 0
    count := 0
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            if grid[i][j] == 0 {
                continue
            } else {
                dfs(i, j, &grid, &count, &result)
                count = 0
            }
        }
    }
    return result
}

func dfs(i, j int, grid *[][]int, count, maxArea *int) {
    if i<0 || j<0 || i>=len((*grid)) || j>=len((*grid)[i]) || (*grid)[i][j]==0 {
        return
    }
    if (*grid)[i][j] == 1 {
        (*grid)[i][j] = 0
        *count++
    }
    if *count > *maxArea {
        *maxArea = *count
    }
    dfs(i-1, j, grid, count, maxArea)
    dfs(i+1, j, grid, count, maxArea)
    dfs(i, j-1, grid, count, maxArea)
    dfs(i, j+1, grid, count, maxArea)

}