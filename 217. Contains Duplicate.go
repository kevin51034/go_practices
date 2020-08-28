func containsDuplicate(nums []int) bool {
    check := make(map[int]bool)
    for _,v := range nums {
        if check[v] == true {
            return true
        }
        check[v] = true
    }
    return false
}