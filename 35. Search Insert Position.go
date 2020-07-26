// Time complexity: O(n)
// Space complexity: O(1)

func searchInsert(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    
    for left<right {
        mid := (left + right)/2
        if target < nums[mid] {
            right = mid - 1
        } else if target > nums[mid] {
            left = mid + 1
        } else if nums[mid] == target{
            return mid
        }
    }
    if target <= nums[left] {
        return left
    }
    return left + 1
}