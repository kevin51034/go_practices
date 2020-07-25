// Time complexity: O(n)
// Space complexity: O(1)

func singleNumber(nums []int) int {
    ones := 0
    twos := 0
    for i:=0; i<len(nums); i++ {
        ones = (ones^nums[i]) &^ twos
        twos = (twos^nums[i]) &^ ones
    }
    return ones
}