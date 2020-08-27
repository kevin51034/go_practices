func intersect(nums1 []int, nums2 []int) []int {
    target := make(map[int]int)

    result := make([]int, 0)
    for _,v := range nums1 {
        target[v]++
    }
    for _,v := range nums2 {
        if target[v] > 0 {
            result = append(result, v)
            target[v]--
        }
    }
    return result
}