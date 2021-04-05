// Time complexity: O(logn)
// Space complexity: O(1)

func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}