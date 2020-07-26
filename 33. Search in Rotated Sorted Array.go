func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    left := 0
    right := len(nums)-1
    
    for left<right {
        mid := (left + right)/2
        if nums[mid] > nums[right] {
            left = left + 1
        } else {
            right = mid
        }
    }
    
    pivot := left
    left = 0
    right = len(nums)-1
    if nums[pivot] <= target && nums[right] >= target {
        left = pivot
    } else {
        right = pivot - 1
    }
    for left<=right {
        mid := (left + right)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}