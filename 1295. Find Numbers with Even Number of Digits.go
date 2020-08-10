func findNumbers(nums []int) int {
    result := 0
    for i:=0; i<len(nums); i++ {
        if 9<nums[i] && nums[i]<100 {
            result++
        } else if 999<nums[i] && nums[i]<10000 {
            result++
        } else if nums[i] == 100000 {
            result++
        }
    }
    return result
}