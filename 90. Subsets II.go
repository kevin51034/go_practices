func subsetsWithDup(nums []int) [][]int {
    result := make([][]int, 0)
    set := make([]int, 0)
    sort.Ints(nums)
    for i:=0; i<=len(nums); i++ {
        findSubsets(nums, i, 0, set, &result)
    }
    return result
}

func findSubsets(nums []int, length, index int, set []int, result *[][]int) {
    if len(set) == length {
        tmp := make([]int, len(set))
        copy(tmp, set)
        *result = append(*result, tmp)
        return
    }
    for i:=index; i<len(nums); i++{
        if i>index && nums[i] == nums[i-1] {
            continue
        }
        num := nums[i]
        set = append(set, num)
        findSubsets(nums, length, i+1, set, result)
        set = set[:len(set)-1]
    }
    return
}