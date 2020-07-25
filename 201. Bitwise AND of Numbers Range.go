// Time complexity: O(n)
// Space complexity: O(1)

// solution1

func rangeBitwiseAnd(m int, n int) int {
    count := 0
    for m!=n {
        m >>= 1
        n >>= 1
        count++
    }
    return m << count
}

// solution2 Brian Kernighan's algorithm

func rangeBitwiseAnd(m int, n int) int {
    for m<n {
        n = n&(n-1)
    }
    return n
}