// Time complexity: O(n)
// Space complexity: O(n)

func subarraySum(nums []int, k int) int {
    result := 0
    sum := 0
    m := make(map[int]int, len(nums)+1)
    m[0] = 1
    for i:=0; i<len(nums); i++ {
        sum+=nums[i]
        if val,ok := m[sum-k]; ok {
            result += val
        }
        m[sum]++
    }
    return result
}