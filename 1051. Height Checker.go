func heightChecker(heights []int) int {
    check := make([]int, len(heights))
    copy(check, heights)
    sort.Ints(check)
    count := 0
    for i:=0; i<len(heights); i++ {
        if check[i] != heights[i] {
            count++
        }
    }
    return count
}