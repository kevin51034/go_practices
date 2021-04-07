// using top-down dp
// Time complexity: O(n^2)
// Space complexity: O(n)

func canJump(nums []int) bool {
    f := make([]bool, len(nums))
    if len(nums)<=1 {
        return true
    }
    f[0] = true
    for i:=0; i<len(nums)-1; i++ {
        if f[i] == false {
            continue
        }
        for j:=0 ;j<=nums[i]; j++ {
            if i+j < len(nums) {
                f[i+j] = true
            }
        }
    }
    return f[len(nums)-1]
}

// greedy
// Time complexity: O(n)
// Space complexity: O(1)

func canJump(nums []int) bool {
    if len(nums) <= 1 {
        return true
    }
    last := len(nums)-1
    for i:=len(nums)-1; i>=0; i-- {
        if i+nums[i] >= last {
            last = i
        }
    }
    return last == 0
}