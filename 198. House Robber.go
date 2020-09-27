// Time complexity: O(n)
// Space complexity: O(n)

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dp := make([]int, len(nums))
    result := 0
    for i:=0; i<len(nums); i++ {
        if i==0 {
            dp[i] = nums[i]
        } else if i==1{
            dp[i] = max(dp[i-1], nums[i])
        } else {
            dp[i] = max((dp[i-2] + nums[i]), dp[i-1])
        }
        if dp[i] > result {
            result = dp[i]
        }
    }
    fmt.Println(dp)
    return result
}

func max(x, y int) int {
    if x>y{
        return x
    }
    return y
}

// optimize
// Time complexity: O(n)
// Space complexity: O(1)

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    curMax, preMax := 0, 0
    for i:=0; i<len(nums); i++ {
        tmp := curMax
        curMax = max(curMax, preMax+nums[i])
        preMax = tmp
    }
    return curMax
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}