func validMountainArray(A []int) bool {
    if len(A) < 3 {
        return false
    }
    toggle := 0
    for i:=0; i<len(A)-1; i++ {
        if A[i] == A[i+1] {
            return false
        }
        if toggle == 0 {
            if A[i] > A[i+1] {
                if i==0 {
                    return false
                }
                toggle = 1
            }
        } else {
            if A[i] < A[i+1] {
                return false
            }
        }
    }
    if toggle == 0 {
        return false
    }
    return true
}