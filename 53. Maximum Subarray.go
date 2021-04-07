// DP
// Time complexity: O(n)
// Space complexity: O(n)

func maxSubArray(nums []int) int {
    dp := make([]int, len(nums))
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    result := nums[0]
    for i:=0; i<len(nums); i++ {
        if i == 0 {
            dp[0] = nums[0]
            continue
        }
        if dp[i-1] > 0 {
            dp[i] = dp[i-1] + nums[i]
        } else {
            dp[i] = nums[i]
        }
        if dp[i] > result {
            result = dp[i]
        }
    }
    return result
}

// optimize
// Time complexity: O(n)
// Space complexity: O(1)
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    result := nums[0]
    preSum := nums[0]
    for i:=1; i<len(nums); i++ {
        if preSum < 0 {
            preSum = nums[i]
        } else {
            preSum += nums[i]
        }
        result = max(result, preSum)
    }
    return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}