func countPrimes(n int) int {
    if n <= 1 {
        return 0
    }
    notPrime := make(map[int]bool, n)
    result := 0
    for i:=2; i<n; i++ {
        if notPrime[i] == false{
            result++
            for j:=i; j<=n; j+=i {
                notPrime[j] = true
            }
        }
    }
    return result
}