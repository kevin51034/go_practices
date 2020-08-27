func twoSum(nums []int, target int) []int {
    another := make(map[int]int)
    for i,v := range nums {
        if _,ok := another[target-v]; ok {
            return []int{another[target-v], i}
        }
        another[v] = i
    }
    return nil
}