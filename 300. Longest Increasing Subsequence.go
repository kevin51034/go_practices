// Time complexity: O(n^2)
// Space complexity: O(n)

func lengthOfLIS(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    f := make([]int, len(nums))
    for i:=0; i<len(nums); i++ {
        f[i] = 1
        for j:=0; j<i; j++ {
            if nums[i] > nums[j] {
                f[i] = max(f[i], f[j]+1)
            }
        }
    }
    return maxElement(f)
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}

func maxElement(a []int) int {
    max := math.MinInt64
    for _,v := range a {
        if v>max {
            max = v
        }
    }
    return max
}