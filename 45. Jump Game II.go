// using top-down dp
// Time complexity: O(n^2)
// Space complexity: O(n)

func jump(nums []int) int {
    f := make([]int, len(nums))
    f[0] = 0
    for i:=0; i<len(nums); i++ {
        if f[i]==0  && i!=0  || nums[i]==0 {
            continue
        }
        for j:=1; j<=nums[i]; j++ {
            if i+j >= len(nums) {
                continue
            }
            if f[i+j] == 0 {
                f[i+j] = f[i]+1
                continue
            }
            f[i+j] = min(f[i]+1, f[i+j])
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
    curFar := 0
    curEnd := 0
    jump := 0
    for i:=0; i<len(nums)-1; i++ {
        curFar = max(curFar, i + nums[i])
        if i==curEnd {
            jump += 1
            curEnd = curFar
        }
    }
    return jump
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}