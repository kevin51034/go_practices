// time complexity: O(nlogn)
// space complexity: O(1)

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    for i:=0; i<len(matrix); i++ {
        left, right := 0, len(matrix[i])-1
        for left<=right {
            mid := left + (right - left)/2
            if matrix[i][mid] == target {
                return true
            } else if matrix[i][mid] > target {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
    }
    return false
}

// time complexity: O(n+m)
// space complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    for i,j:=0, len(matrix[0])-1; i<len(matrix) && j>=0; {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] < target {
            i++
        } else {
            j--
        }
    }
    return false
}