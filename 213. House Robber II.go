// Time complexity: O(n)
// Space complexity: O(1)

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    } else if len(nums) == 1 {
        return nums[0]
    }
    result := 0
    curSum, preSum := 0, 0
    for i:=0; i<len(nums)-1; i++ {
        tmp := curSum
        curSum = max(curSum, preSum + nums[i])
        preSum = tmp
    }
    result = curSum
    curSum, preSum = 0, 0
    for i:=1; i<len(nums); i++ {
        tmp := curSum
        curSum = max(curSum, preSum + nums[i])
        preSum = tmp
    }
    return max(result, curSum)
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}