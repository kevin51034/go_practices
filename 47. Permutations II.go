func permuteUnique(nums []int) [][]int {
    result := make([][]int, 0)
    set := make([]int, 0)
    used := make([]bool, len(nums))
    sort.Ints(nums)
    findPermutations(nums, set, &used, &result)
    return result
}

func findPermutations(nums, set []int, used *[]bool, result *[][]int) {
    if len(set) == len(nums) {
        tmp := make([]int, len(set))
        copy(tmp, set)
        *result = append(*result, tmp)
        return
    }
    for i:=0; i<len(nums); i++ {
        if (*used)[i] {
            continue
		}
		// Same number can be only used once at each depth.
        if i>0 && nums[i] == nums[i-1] && !(*used)[i-1] {
            continue
        }
        (*used)[i] = true
        set = append(set, nums[i])
        findPermutations(nums, set, used, result)
        set = set[:len(set)-1]
        (*used)[i] = false
    }
    return
}
