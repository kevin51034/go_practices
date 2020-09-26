// Time complexity: O(mn)
// Space complexity: O(1) // no extra space

func minFallingPathSum(A [][]int) int {
    for i:=0; i<len(A); i++ {
        for j:=0; j<len(A[i]); j++ {
            if i==0 {
                continue
            }
            if j==0 {
                A[i][j] = min(A[i-1][j], A[i-1][j+1]) + A[i][j]
            } else if j==len(A[i])-1 {
                A[i][j] = min(A[i-1][j], A[i-1][j-1]) + A[i][j]
            } else {
                A[i][j] = min(min(A[i-1][j-1], A[i-1][j]), A[i-1][j+1]) + A[i][j]
            }
        }
    }
    return minElement(A[len(A)-1])
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

func minElement(A []int) int {
    min := math.MaxInt64
    for _,num := range A {
        if num < min {
            min = num
        }
    }
    return min
}