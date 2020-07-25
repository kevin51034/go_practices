// using xor operation
// Time complexity: O(n)
// Space complexity: O(1)

func singleNumber(nums []int) int {
    result := 0
    for i:=0; i<len(nums);i++ {
        result ^= nums[i]
    }
    return result
}