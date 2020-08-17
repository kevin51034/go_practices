func removeDuplicates(nums []int) int {
    if len(nums)<=1 {
        return len(nums)
    }
    count:= 1
    tmp := nums[0]
    for i, j:=1,1; i<len(nums); i++ {
        if nums[i] == tmp {
            continue
        } else {
            tmp = nums[i]
            nums[i] , nums[j] = nums[j], nums[i]
            j++
            count++
        }
    }
    return count
}