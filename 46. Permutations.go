// Time complexity: O(n!)
// Space complexity: O(n)
func permute(nums []int) [][]int {
    result := make([][]int, 0)
    set := make([]int, 0)
    used := make([]bool, len(nums))
    findPermutations(nums, set, &used, &result)
    return result
}

func findPermutations(nums []int, set[]int, used*[]bool, result*[][]int) {
    if len(set) == len(nums) {
        tmp := make([]int, len(set))
        copy(tmp, set)
        *result = append(*result, tmp)
        return
    }
    for i:=0; i<len(nums); i++ {
        num := nums[i]
        if (*used)[i] {
            continue
        }
        (*used)[i] = true
        set = append(set, num)
        findPermutations(nums, set, used, result)
        set = set[:len(set)-1]
        (*used)[i] = false
    }
    return
}