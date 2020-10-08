// time complexity: O(n)
// space complexity: O(1)

func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    result := nums[0]
    maxSum, minSum := nums[0], nums[0]
    for i:=1; i<len(nums); i++ {
        if nums[i]<0 {
            maxSum, minSum = minSum, maxSum
        }
        maxSum = max(nums[i], maxSum*nums[i])
        minSum = min(nums[i], minSum*nums[i])
        result = max(result, maxSum)
    }
    return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}