// solution 1
// sort

func sortedSquares(A []int) []int {
    result := make([]int, 0)
    for i:=0; i<len(A); i++ {
        result = append(result, A[i]*A[i])
    }
    sort.Sort(sort.IntSlice(result))
    return result
}

// solution 2
// two pointer

func sortedSquares(A []int) []int {
    result := make([]int, len(A))
    
    for left, right, i:=0, len(A)-1, len(A)-1;i>=0; i-- {
        if A[left]*A[left] > A[right]*A[right] {
            result[i] = A[left]*A[left]
            left++
        } else {
            result[i] = A[right]*A[right]
            right--
        }
    }
    return result
}