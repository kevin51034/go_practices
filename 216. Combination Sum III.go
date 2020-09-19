func combinationSum3(k int, n int) [][]int {
    result := make([][]int, 0)
    set := make([]int, 0)
    findCombinations(k, n, 1, set, &result)
    return result
}

func findCombinations(k, target, index int, set []int, result *[][]int) {
    if len(set) == k && target == 0 {
        tmp := make([]int, len(set))
        copy(tmp, set)
        *result = append(*result, tmp)
        return
    } else if target < 0 || len(set)>k {
        return
    }
    for i:=index; i<10; i++ {
        set = append(set, i)
        findCombinations(k, target-i, i+1, set, result)
        set = set[:len(set)-1]
    }
    return
}