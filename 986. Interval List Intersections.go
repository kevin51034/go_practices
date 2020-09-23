// Time complexity: O(m + n)
// Space complexity: O(1)

func intervalIntersection(A [][]int, B [][]int) [][]int {
    result := make([][]int, 0)
    for i,j:=0,0; i<len(A) && j <len(B); {
        start := max(A[i][0], B[j][0])
        end := min(A[i][1], B[j][1])
        if start<=end {
            result = append(result, []int{start,end})
        }
        if A[i][1] <= B[j][1] {
            i++
        } else {
            j++
        }
    }
    return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}