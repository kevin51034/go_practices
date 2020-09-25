// Time complexity: O(n)
// Space complexity: O(n)

func longestSubsequence(arr []int, difference int) int {
    m := map[int]int{} 
    result := 0
    for _,num := range arr {
        m[num] = m[num - difference] + 1
        if m[num] > result {
            result = m[num]
        }
    }
    return result
}