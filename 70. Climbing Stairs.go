// Time complexity: O(n)
// Space complexity: O(n)

func climbStairs(n int) int {
    f := make([]int, n+1)
    f[0], f[1] = 1,1
    for i:=2; i<=n; i++ {
            f[i] = f[i-1] + f[i-2]
    }
    return f[n]
}


// optimize
// Time complexity: O(n)
// Space complexity: O(1)

func climbStairs(n int) int {
    f0, f1:= 1, 1
    if n==0 || n==1 {
        return 1
    }
    for i:=2; i<n; i++ {
        tmp := f0 + f1
        f0 = f1
        f1 = tmp
    }
    return f0 + f1
}