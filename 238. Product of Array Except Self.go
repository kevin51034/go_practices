// time complexity: O(n)
// space complexity: O(n)

func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))
    result[0] = 1
    result[len(nums)-1] = 1
    for i:=1; i<len(nums); i++ {
        result[i] = nums[i-1]*result[i-1]
    }
    right := 1
    for i:=len(nums)-2; i>=0; i-- {
        right*=nums[i+1]
        result[i] = right*result[i]
    }
    return result
}