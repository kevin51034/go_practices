func findDisappearedNumbers(nums []int) []int {
    result := make([]int, 0)
    for i:=0; i<len(nums); i++ {
        if nums[abs(nums[i])-1] > 0 {
            nums[abs(nums[i])-1] *= -1
        }
    }
    for i:=0; i<len(nums); i++ {
        if nums[i] > 0 {
            result = append(result, i+1)
        }
    }
    return result
}

func abs(x int) int {
    if x>0 {
        return x
    }
    return -x
}

