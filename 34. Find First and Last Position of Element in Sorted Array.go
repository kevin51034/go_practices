func searchRange(nums []int, target int) []int {
    left, right := 0, len(nums)-1
	result := make([]int, 0)
	// find first position
    for left <= right {
        mid := left + (right - left)/2
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            if mid == 0 || nums[mid-1] != target {
                result = append(result, mid)
                break
            } else {
                right = mid - 1
            }
        }
    }
	left, right = 0, len(nums)-1
	// find last position
    for left <= right {
        mid := left + (right - left)/2
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            if mid == len(nums)-1 || nums[mid+1] != target {
                result = append(result, mid)
                break
            } else {
                left = mid + 1
            }
        }
    }
    if len(result) == 0 {
        result = []int{-1,-1}
    }
    return result
}