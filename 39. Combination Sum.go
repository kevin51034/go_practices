// Time compleixty: O(2^n)
// Space complexity: O(n)

func combinationSum(candidates []int, target int) [][]int {
    result := make([][]int, 0)
    if len(candidates) == 0 {
        return result
    }
    oneset := make([]int, 0)
    findCombinations(candidates, target, 0, &oneset, &result)
    return result
}

func findCombinations(candidates []int, target, index int, oneset *[]int, result *[][]int) {
    if target < 0 {
        return
    }
    if target == 0 {
        tmp := make([]int, len(*oneset))
        copy(tmp, *oneset)
        *result = append(*result, tmp)
        return
    }
    for i:=index; i<len(candidates); i++ {
        num := candidates[i]
        *oneset = append(*oneset, num)
        findCombinations(candidates, target-num, i, oneset, result)
        *oneset = (*oneset)[:len(*oneset)-1]
    }
    return
}