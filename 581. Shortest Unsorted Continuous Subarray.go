// Time complexity: O(n)
// Space complexity: O(1)

func findUnsortedSubarray(nums []int) int {
    max, min := math.MinInt64, math.MaxInt64
    left, right := -1, -1
    for i:=0; i<len(nums); i++ {
        if nums[i] >= max {
            max = nums[i]
        } else {
            right = i
        }
    }
    if right == -1 {
        return 0
    }
    for i:=len(nums)-1; i>=0; i-- {
        if nums[i] <= min {
            min = nums[i]
        } else {
            left = i
        }
    }
    return right - left + 1
}