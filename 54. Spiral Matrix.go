// Time complexity: O(n)
// Space complexity: O(1)
// be caution to the boundary

func spiralOrder(matrix [][]int) []int {
    result := make([]int, 0)
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return result
    }
    // set the bound
    top := 0
    bottom := len(matrix)-1
    left := 0
    right := len(matrix[0])-1
    length := len(matrix) * len(matrix[0])
    for len(result) < length {
        for i:=left; i<=right&&len(result)<length; i++ {
            result = append(result, matrix[top][i])
        }
        top++
        
        for i:=top; i<=bottom&&len(result)<length; i++ {
            result = append(result, matrix[i][right])
        }
        right--
        
        for i:=right; i>=left&&len(result)<length; i-- {
            result = append(result, matrix[bottom][i])
        }
        bottom--
        
        for i:=bottom; i>=top&&len(result)<length; i-- {
            result = append(result, matrix[i][left])
        }
        left++
    }
    return result
}