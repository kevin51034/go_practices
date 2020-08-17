func sortArrayByParity(A []int) []int {
    for i,j:=0,0; i<len(A);i++{
        if A[i]%2 == 0 {
            A[i], A[j] = A[j], A[i]
            j++
        }
    }
    return A
}