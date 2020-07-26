func searchMatrix(matrix [][]int, target int) bool {
    left := 0
    right := len(matrix)-1
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    for left<=right {
        mid := (left + right)/2
        if target <= matrix[mid][len(matrix[mid])-1] && target >= matrix[mid][0] {
            return searchInRows(matrix[mid], target)
        } else if target < matrix[mid][0] {
            right = mid - 1
        } else {
            left = mid + 1
        } 
    }
    return false
}

func searchInRows(row[]int, target int) bool {
    left := 0
    right := len(row)-1
    for left<=right {
        mid:=(left + right)/2
        if target < row[mid] {
            right = mid - 1
        } else if target > row[mid] {
            left = mid + 1
        } else if target == row[mid] {
            return true
        }
    }
    return false
}