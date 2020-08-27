func intersection(nums1 []int, nums2 []int) []int {
    check := make(map[int]bool)
    result := make([]int, 0)
    for _,v := range nums1 {
        check[v] = true
    }
    for _,v := range nums2 {
        if check[v] == true {
            check[v] = false
            result = append(result, v)
        }
    }
    return result
}