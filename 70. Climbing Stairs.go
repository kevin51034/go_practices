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

// with memoization
// Time complexity: O(n)
// Space complexity: O(n)

func climbStairs(n int) int {
    memo := make([]int, n+1)
    return helper(n, memo)
}

func helper(n int, memo[]int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    if memo[n] != 0 {
        return memo[n]
    }
    memo[n] = helper(n-1, memo) + helper(n-2, memo)
    return memo[n]
}