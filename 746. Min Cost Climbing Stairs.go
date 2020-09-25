// DP
// Time complexity: O(n)
// Space complexity: O(1)

func minCostClimbingStairs(cost []int) int {
    if len(cost) <= 1 {
        return 0
    }
    for i:=2; i<len(cost); i++ {
        cost[i] = min(cost[i-1], cost[i-2]) + cost[i]
    }
    return min(cost[len(cost)-2], cost[len(cost)-1])
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

// Recursion + Memorization
// Time complexity: O(n)
// Space complexity: O(n)

func minCostClimbingStairs(cost []int) int {
    if len(cost) <= 1 {
        return 0
    }
    m := make([]int, len(cost))
    dp(cost, m, len(cost)-1)
    return min(m[len(m)-1], m[len(m)-2])
}

func dp(cost, m []int, index int) int {
    if index < 0 {
        return 0
    }
    if m[index] > 0 {
        return m[index]
    }
    m[index] = min(dp(cost, m, index-1), dp(cost, m, index-2)) + cost[index]
    return m[index]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}