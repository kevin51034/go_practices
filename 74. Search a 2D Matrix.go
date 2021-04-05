// solution 1

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    left, right := 0, len(matrix)-1
    for left <= right {
        mid := left + (right - left)/2
        if matrix[mid][0] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if left == 0 {
        return false
    }
    left--
    for i:=0; i<len(matrix[left]); i++ {
        if matrix[left][i] == target {
            return true
        }
    }
    return false
}

// solution 2
// binary search two times
// first search the row of the matrix, and then search the column of the target
// be careful that if target doesn't exist, left pointer might cross the boundary
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    left, right := 0, len(matrix)-1
    for left <= right {
        mid := left + (right - left)/2
        if matrix[mid][0] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if left == 0 {
        return false
    }
    index := left-1
    left, right = 0, len(matrix[index])-1
    for left <= right {
        mid := left + (right - left)/2
        if matrix[index][mid] >= target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if left >= len(matrix[index]) {
        return false
    }
    if matrix[index][left] == target {
        return true
    }
    return false
}