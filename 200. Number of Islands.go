// Time complexity: O(nm)
// Space complexity: O(nm)

// grid doesn't have to pass in the pointer because of slice
// see
// https://kitecloud-backend.coderbridge.io/2020/08/15/golang-slice-%E4%BD%9C%E7%82%BA%E5%8F%83%E6%95%B8%E5%82%B3%E9%81%9E%E6%99%82%E7%9A%84%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A0%85/

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    result := 0
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            if grid[i][j] == '0' {
                continue
            } else {
                result++
                dfs(i, j, grid)
            }
        }
    } 
    return result
}

func dfs(i, j int, grid [][]byte) {
    if i<0 || j<0 || i>=len(grid) || j>=len(grid[i]) || grid[i][j] == '0' {
        return
    }
    grid[i][j] = '0'
    dfs(i-1, j, grid)
    dfs(i+1, j, grid)
    dfs(i, j-1, grid)
    dfs(i, j+1, grid)
    return
}