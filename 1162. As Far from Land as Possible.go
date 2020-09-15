// Time compleixty: O(n^2)
// Space complexity: O(n^2)

func maxDistance(grid [][]int) int {
    queue := [][]int{}
    dist := make([][]int, len(grid))
    ok := false
    
    for i:=0; i<len(grid); i++ {
        dist[i] = make([]int, len(grid[i]))
    }
    
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            if grid[i][j] == 1 {
                queue = append(queue, []int{i,j})
            } else {
                ok = true
            }
        }
    }
    
    if ok == false {
        return -1
    }
    step := 0
    for len(queue) > 0 {
        step++
        for l:=len(queue)-1; l>=0; l-- {
            node := queue[0]
            queue = queue[1:]
            for _,diff := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
                if i,j:=node[0]+diff[0], node[1]+diff[1]; i>=0 && i<len(grid) && j>=0 && j<len(grid[i]) {
                    if grid[i][j] != 0 {
                        continue
                    }
                    if dist[i][j] == 0 {
                        dist[i][j] = step
                        queue = append(queue, []int{i, j})
                    }
                }
            }
        }
    }
    return step-1
}