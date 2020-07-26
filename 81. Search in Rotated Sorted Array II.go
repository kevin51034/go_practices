func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    
    left := 0
    right := len(nums) - 1
    for left < right {
        for left < right && nums[left] == nums[left+1] {
            left = left + 1
        }
        for left < right && nums[right] == nums[right-1] {
            right = right -1
        }
        mid := (left + right)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    pivot := left
    left = 0
    right = len(nums) - 1
    if target >= nums[pivot] && target <= nums[right] {
        left = pivot
    } else {
        right = pivot - 1
    }
    
    for left <= right {
        mid := (left + right)/2
        if nums[mid] == target {
            return true
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}