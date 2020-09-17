func combine(n int, k int) [][]int {
    result := make([][]int, 0)
    set := make([]int, 0)
    findCombinations(n, k, 1, set, &result)
    return result
}

func findCombinations(n, k, index int, set []int, result *[][]int) {
    if len(set) > k {
        return
    }
    if len(set) == k {
        tmp := make([]int, len(set))
        copy(tmp, set)
        *result =append(*result, tmp)
        return
	}
	// i will at most be n - (k - set.size()) + 1
    for i:=index; i<=n-(k-len(set))+1; i++ {
        set = append(set, i)
        findCombinations(n, k, i+1, set, result)
        set = set[:len(set)-1]
    }
    return
}