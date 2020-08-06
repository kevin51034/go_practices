func numTrees(n int) int {
    f := make([]int, n+1)
    f[0] = 1
    f[1] = 1
    for i:=2; i<=n; i++ {
        for j:=1; j<=i; j++ {
            f[i] += f[j-1] * f[i-j]
        }
    }
    return f[n]
}