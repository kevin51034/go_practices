// Time complexity: O(nm) n, m = rows, columns
// Space complexity: O(nm)

func updateMatrix(matrix [][]int) [][]int {
    queue := make([][]int, 0)
    
    for i:=0; i<len(matrix); i++ {
        for j:=0; j<len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                queue = append(queue, []int{i,j})
            } else {
                matrix[i][j] = -1
            }
        }
    }
    neighbors := [][]int{{0,1},{0,-1},{-1,0},{1,0}}
    for len(queue)!=0 {
        originPoint := queue[0]
        queue = queue[1:]
        
        for _,v:= range neighbors {
            x := originPoint[0] + v[0]
            y := originPoint[1] + v[1]
            if x>=0 && x<len(matrix) && y>=0 && y<len(matrix[x]) && matrix[x][y]==-1 {
                matrix[x][y] = matrix[originPoint[0]][originPoint[1]] + 1
                queue = append(queue, []int{x,y})
            }
        }
    }
    return matrix
}