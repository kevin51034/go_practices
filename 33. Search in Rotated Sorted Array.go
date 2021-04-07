func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right - left)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    pivot := left
    if nums[pivot] <= target && nums[len(nums)-1] >= target {
        right = len(nums)-1
    } else {
        left = 0
        right = pivot - 1
    }
    for left <= right {
        mid := left + (right - left)/2
        if nums[mid] >= target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if nums[left] == target {
        return left
    }
    return -1
}