// using top-down dp
// Time complexity: O(n^2)
// Space complexity: O(n)

func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    f := make([]int, len(nums))

    for i:=0; i<len(nums); i++ {
        for j:=1; j<=nums[i] && i+j<len(nums); j++ {
            if f[i+j] == 0 {
                f[i+j] = f[i]+1
                continue
            }
            f[i+j] = min(f[i+j], f[i]+1)
        }
    }
    return f[len(nums)-1]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

// greedy
// Time complexity: O(n)
// Space complexity: O(1)

func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    curMax := 0
    curEnd := 0
    result := 0
    for i:=0; i<len(nums); i++ {
        if i > curEnd {
            result ++
            curEnd = curMax
        }
        curMax = max(curMax, i+nums[i])
    }
    return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}