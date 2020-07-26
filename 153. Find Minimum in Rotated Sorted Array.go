// Time complexity: O(logn)
// Space complexity: O(n)

func findMin(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    left := 0
    right := len(nums)-1
    
    for left<right {
        mid := (left + right)/2
        if nums[mid] < nums[right] {
            right = mid
        } else {
            left = mid + 1
        }
    }
    if nums[left] > nums[right] {
        return nums[right]
    }
    return nums[left]
}