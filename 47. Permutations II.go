// Time complexity: O(n!)
// Space complexity: O(n + k)
func permuteUnique(nums []int) [][]int {
    result := make([][]int, 0)
    set := make([]int, 0)
    used := make([]bool, len(nums))
    sort.Ints(nums)
    findPermutations(nums, used, set, &result)
    return result
}

func findPermutations(nums []int, used []bool, set []int, result *[][]int) {
    if len(nums) == len(set) {
        tmp := make([]int, len(set))
        copy(tmp, set)
        *result = append(*result, tmp)
        return
    }
    for i:=0; i<len(nums); i++ {
        if used[i] {
            continue
        }
		// Same number can only be used once at each depth.
        if i>0 && nums[i] == nums[i-1] && !used[i-1] {
            continue
        }
        used[i] = true
        set = append(set, nums[i])
        findPermutations(nums, used, set, result)
        set = set[:len(set)-1]
        used[i] = false
    }
    return
}