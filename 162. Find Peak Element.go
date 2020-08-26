func findPeakElement(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[mid+1] {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func findPeakElement(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right - left)/2
        if mid == len(nums)-1 {
            return mid
        }
        if nums[mid] > nums[mid+1] {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}