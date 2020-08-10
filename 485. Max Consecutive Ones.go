func findMaxConsecutiveOnes(nums []int) int {
    result := 0
    tmp := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] == 1 {
            tmp++
            result = max(result, tmp)
        } else {
            tmp = 0
        }
    }
    return result
}

func max(x, y int) int{
    if x>y {
        return x
    }
    return y
}